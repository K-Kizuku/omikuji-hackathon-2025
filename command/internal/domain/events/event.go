package events

import (
	"github.com/K-Kizuku/pymon-command-api/internal/domain/models"
	esa "github.com/j5ik2o/event-store-adapter-go/pkg"
)

type EventMetadata struct {
	id         string
	seqNr      uint64
	executorId models.UserAccountId
	occurredAt uint64
}

type GroupChatEvent interface {
	esa.Event
	GetExecutorId() *models.UserAccountId
	ToProto() []byte
}
