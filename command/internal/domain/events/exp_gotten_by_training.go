package events

import (
	"fmt"
	"time"

	"github.com/K-Kizuku/pymon-command-api/internal/domain/models"
	"github.com/K-Kizuku/pymon-command-api/pkg/protobuf"
	"github.com/K-Kizuku/pymon-command-api/pkg/protobuf/pymon"
	"github.com/google/uuid"
	esa "github.com/j5ik2o/event-store-adapter-go/pkg"
	"google.golang.org/protobuf/proto"
)

type ExpGottenByTraining struct {
	EventMetadata
	aggregateID models.PymonID
	pymonID     string
	exp         int64
}

func NewExpGottenByTraining(aggregateID models.PymonID, pymonID string, exp int64, seqNr uint64, executorId models.UserAccountID) ExpGottenByTraining {
	id := uuid.Must(uuid.NewV7())
	now := time.Now()
	occurredAt := uint64(now.UnixNano() / 1e6)
	return ExpGottenByTraining{
		EventMetadata: EventMetadata{
			id:         id.String(),
			seqNr:      seqNr,
			executorId: executorId,
			occurredAt: occurredAt,
		},
		aggregateID: aggregateID,
		pymonID:     pymonID,
		exp:         exp,
	}
}

func NewExpGottenByTrainingFrom(id string, aggregateID models.PymonID, pymonID string, exp int64, seqNr uint64, executorId models.UserAccountID, occurredAt uint64) ExpGottenByTraining {
	return ExpGottenByTraining{
		EventMetadata: EventMetadata{
			id:         id,
			seqNr:      seqNr,
			executorId: executorId,
			occurredAt: occurredAt,
		},
		aggregateID: aggregateID,
		pymonID:     pymonID,
		exp:         exp,
	}
}

func (e *ExpGottenByTraining) ToProto() []byte {

	p := &protobuf.EventEnvelope{
		Metadata: &protobuf.EventMetadata{
			Id:          e.id,
			AggregateId: e.aggregateID.GetValue(),
			SeqNr:       e.seqNr,
			OccurredAt:  e.occurredAt,
			ExecutorId:  e.executorId.GetValue(),
		},
		Event: &protobuf.EventEnvelope_ExpGottenByTraining{
			ExpGottenByTraining: &pymon.ExpGottenByTraining{
				PymonId: e.pymonID,
				Exp:     e.exp,
			},
		},
	}
	b, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	return b
}

func (e *ExpGottenByTraining) GetId() string {
	return e.id
}

func (e *ExpGottenByTraining) GetTypeName() string {
	return "ExpGottenByTraining"
}

func (e *ExpGottenByTraining) GetAggregateId() esa.AggregateId {
	return &e.aggregateID
}

func (e *ExpGottenByTraining) GetSeqNr() uint64 {
	return e.seqNr
}

func (e *ExpGottenByTraining) GetExecutorId() *models.UserAccountID {
	return &e.executorId
}

func (e *ExpGottenByTraining) IsCreated() bool {
	return true
}

func (e *ExpGottenByTraining) GetOccurredAt() uint64 {
	return e.occurredAt
}

func (e *ExpGottenByTraining) String() string {
	return fmt.Sprintf("%s{ id: %s, aggregateId: %s, seqNr: %d, occurredAt: %d}",
		e.GetTypeName(), e.id, e.aggregateID, e.seqNr, e.occurredAt)
}
