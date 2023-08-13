package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
	"google.golang.org/protobuf/proto"

	modelv1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/model/v1"
	"github.com/hampgoodwin/todo/internal/event"
)

func WireTap(url string) (*nats.EncodedConn, error) {
	nc, _ := nats.Connect(url)
	nec, _ := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)

	_, err := nec.Subscribe(">", messageHandler)
	if err != nil {
		return nil, fmt.Errorf("subscribing to all subjects: %w", err)
	}
	return nec, nil
}

func messageHandler(msg *nats.Msg) {
	switch msg.Subject {
	case event.SubjectToDoCreated:
		account := &modelv1.ToDo{}
		err := proto.Unmarshal(msg.Data, account)
		if err != nil {
			fmt.Printf("error unmarshaling message on subject %q", event.SubjectToDoCreated)
		}
		fmt.Printf("received %q\n%v\n", event.SubjectToDoCreated, account)
	default:
		fmt.Printf("unhandled event received on subject %q\n", msg.Sub.Subject)
	}
}
