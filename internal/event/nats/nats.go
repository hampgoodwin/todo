package nats

import (
	"fmt"

	"github.com/hampgoodwin/todo/internal/event"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
)

func NewNATSEncodedConn(url string) (*nats.EncodedConn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("connecting to nats: %w", err)
	}
	nec, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		return nil, fmt.Errorf("creating new encoded connection for protobuf: %w", err)
	}
	return nec, nil
}

func NewNATSJetStream(nec *nats.EncodedConn) (nats.JetStreamContext, error) {
	jsc, err := nec.Conn.JetStream()
	if err != nil {
		return nil, fmt.Errorf("getting jetstream context: %w", err)
	}

	// Configure stream
	todoStreamConfiguration := &nats.StreamConfig{
		Name:        "todo",
		Description: "all todo subjects",
		Subjects:    []string{event.SubjectToDoCreated},
	}
	_, err = jsc.AddStream(todoStreamConfiguration)
	if err != nil {
		return nil, fmt.Errorf("adding account stream: %w", err)
	}

	return jsc, nil
}
