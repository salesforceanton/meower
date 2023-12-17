package web_socket

import "time"

const KindMeowCreated = iota + 1

type MeowCreatedMessage struct {
	Kind      int32     `json:"kind"`
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMeowCreatedMessage(id, body string, createdAt time.Time) *MeowCreatedMessage {
	return &MeowCreatedMessage{
		Kind:      KindMeowCreated,
		ID:        id,
		Body:      body,
		CreatedAt: createdAt,
	}
}
