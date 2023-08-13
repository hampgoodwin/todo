package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/hampgoodwin/todo/internal/database"
	"github.com/hampgoodwin/todo/internal/environment"
	inats "github.com/hampgoodwin/todo/internal/event/nats"
	grpccontroller "github.com/hampgoodwin/todo/internal/grpc/v1/controller"
	grpcrouter "github.com/hampgoodwin/todo/internal/grpc/v1/router"
	"github.com/hampgoodwin/todo/internal/repository"
	"github.com/hampgoodwin/todo/internal/service"
	itrace "github.com/hampgoodwin/todo/internal/trace"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Create the OTLP Tracer Provider
	tpShutdownFn, err := itrace.SetOTLPGRPCTracerProvider(ctx)
	if err != nil {
		log.Panic("failed to create otlp grpc exporter")
	}

	// Load the environment
	// The environment includes the minimum necessary dependencies to start the application
	env, err := environment.New(environment.Environment{}, "/etc/todo/.env.toml")
	if err != nil {
		_ = tpShutdownFn(ctx)
		log.Panic("failed to create new environment")
	}

	// Create the postgres database pool and migrate
	db, err := database.NewDatabasePool(ctx, env.Config.Database.ConnectionString())
	if err != nil {
		env.Log.Error("creating new database pool", zap.Error(err))
		log.Fatal("error creating database pool on application start")
	}
	if err := database.Migrate(db); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			env.Log.Fatal("migrating", zap.Error(err))
			log.Fatal("error migrating database on application start")
		}
		env.Log.Info("no migration changes")
	}

	// Create the repository layer
	repository := repository.NewRepository(db)

	// Create NATS event bus, using proto encoded connection and JetStream
	env.Log.Info("starting nats", zap.Any("service_info", env.Config.NATS.URL()))
	nenc, err := inats.NewNATSEncodedConn(env.Config.NATS.URL())
	if err != nil {
		env.Log.Error("nats error, shutting down", zap.Error(err))
		close(ctx, env.Log, db, nenc, nil, nil, tpShutdownFn)
		log.Fatal("failed to create nats connection")
	}
	env.Log.Info("creating jetstream")
	if _, err := inats.NewNATSJetStream(nenc); err != nil {
		env.Log.Error("jetstream error, shutting down", zap.Error(err))
		close(ctx, env.Log, db, nenc, nil, nil, tpShutdownFn)
		log.Fatal("failed to create jetstream")
	}
	var nencWiretap *nats.EncodedConn
	if env.Config.NATS.Wiretap.Enable {
		env.Log.Info("starting wiretap", zap.Any("service_info", env.Config.NATS.Wiretap.URL()))
		nencWiretap, err = inats.WireTap(env.Config.NATS.Wiretap.URL())
		if err != nil {
			env.Log.Error("nats wiretap error, shutting down", zap.Error(err))
			close(ctx, env.Log, db, nenc, nencWiretap, nil, tpShutdownFn)
			log.Fatal("failed to create wiretap")
		}
	}

	// Create the service layer
	service := service.NewService(env.Log, repository, nenc)

	// Create the gRPC server
	// Create the controller for the to-be-created gRPC Server
	grpcController := grpccontroller.NewController(env.Log, service)
	// Create listener for gRPC Server
	lis, err := net.Listen("tcp", env.Config.GRPCServer.URL())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create the gRPC server with zap log grpc intercepter
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpc_zap.UnaryServerInterceptor(env.Log),
	))
	// Register the controller with the server
	grpcrouter.Register(grpcServer, grpcController)

	// Serve the gRPC server on the earlier created Listener
	grpcErr := make(chan error)
	go func() {
		env.Log.Info("starting grpc server", zap.Any("service_info", grpcServer.GetServiceInfo()))
		grpcErr <- grpcServer.Serve(lis)
	}()

	// Handle any errors from Servers
	for {
		select {
		case err := <-grpcErr:
			env.Log.Error("grpc server error, shutting down", zap.Error(err))
			close(ctx, env.Log, db, nenc, nencWiretap, grpcServer, tpShutdownFn)
			return
		case <-ctx.Done():
			fmt.Printf("received shutdown signal: %s\n", ctx.Err())
			cancel()
			close(ctx, env.Log, db, nenc, nencWiretap, grpcServer, tpShutdownFn)
			return
		}
	}
}

// close cleans up the application dependencies
func close(
	ctx context.Context,
	log *zap.Logger,
	db *pgxpool.Pool,
	nenc *nats.EncodedConn,
	nencWiretap *nats.EncodedConn,
	grpcServer *grpc.Server,
	tpShutdownFunc func(context.Context) error,
) {
	log.Info("closing")
	// close grpc server
	if grpcServer != nil {
		log.Info("closing grpcserver")
		grpcServer.GracefulStop()
	}
	// close http server
	// disconnect from db
	if db != nil {
		log.Info("closing db")
		db.Close()
	}
	// drain nats encoded connection
	if nenc != nil {
		log.Info("draining and closing nats connection")
		if err := nenc.Drain(); err != nil {
			log.Error("draining and closing nats connection", zap.Error(err))
		}
	}
	// drain nats wire tap encoded connection
	if nencWiretap != nil {
		log.Info("draining and closing wiretap connection")
		if err := nencWiretap.Drain(); err != nil {
			log.Error("draining and closing wiretap connection", zap.Error(err))
		}
	}
	// shutdown tracer provider
	log.Info("shutting down tracer provider")
	if err := tpShutdownFunc(ctx); err != nil {
		log.Error("shutting down tracer provider", zap.Error(err))
	}
}
