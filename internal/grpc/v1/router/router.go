package router

import (
	servicev1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/service/v1"
	"github.com/hampgoodwin/todo/internal/grpc/v1/controller"
	"google.golang.org/grpc"
)

func Register(
	srv *grpc.Server,
	ctrl *controller.Controller,
) {
	servicev1.RegisterToDoServiceServer(srv, ctrl)
}
