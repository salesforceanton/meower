package eventbus

import (
	"bytes"
	"encoding/gob"

	"github.com/nats-io/nats.go"
	"github.com/salesforceanton/meower/internal/schema"
)

type NatsEventbus struct {
	connect                 *nats.Conn
	meowCreatedSubscription *nats.Subscription
}

func NewNatsEventbus(addr string) (*NatsEventbus, error) {
	connect, err := nats.Connect(addr)
	if err != nil {
		return nil, err
	}
	return &NatsEventbus{connect: connect}, nil
}

func (e *NatsEventbus) Close() {
	if e.meowCreatedSubscription != nil {
		e.meowCreatedSubscription.Unsubscribe()
	}
	e.connect.Close()
}

func (e *NatsEventbus) PublishMeowCreated(m schema.Meow) error {
	message := MeowCreatedMessage{m.Id, m.Body, m.CreatedAt}
	encodedMessage, err := e.writeMessage(&message)
	if err != nil {
		return err
	}

	return e.connect.Publish(message.Key(), encodedMessage)
}

func (e *NatsEventbus) SubscribeMeowCreated(createHandler func(MeowCreatedMessage)) (err error) {
	message := MeowCreatedMessage{}
	e.meowCreatedSubscription, err = e.connect.Subscribe(message.Key(), func(m *nats.Msg) {
		e.readMessage(m.Data, &message)
		createHandler(message)
	})
	return err
}

func (e *NatsEventbus) writeMessage(m Message) ([]byte, error) {
	bytesWriter := bytes.Buffer{}
	err := gob.NewEncoder(&bytesWriter).Encode(m)
	if err != nil {
		return nil, err
	}
	return bytesWriter.Bytes(), nil
}

func (e *NatsEventbus) readMessage(data []byte, m interface{}) error {
	bytesReader := bytes.Buffer{}
	bytesReader.Write(data)
	return gob.NewDecoder(&bytesReader).Decode(m)
}
