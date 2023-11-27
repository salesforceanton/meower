package eventbus

import "github.com/salesforceanton/meower/internal/schema"

const MEOW_CREATED_MESSAGE_STRING = "meow_created"

type EventBus interface {
	Close()
	PublishMeowCreated(schema.Meow) error
	SubscribeMeowCreated(f func(MeowCreatedMessage)) error
}
