package eventbus

import "time"

type Message interface {
	Key() string
}

type MeowCreatedMessage struct {
	Id        string
	Body      string
	CreatedAt time.Time
}

func (m *MeowCreatedMessage) Key() string {
	return MEOW_CREATED_MESSAGE_STRING
}
