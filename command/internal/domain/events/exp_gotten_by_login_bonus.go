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

type ExpGottenByLoginBonus struct {
	EventMetadata
	aggregateID models.PymonID
	pymonID     string
	exp         int64
}

func NewExpGottenByLoginBonus(aggregateID models.PymonID, pymonID string, exp int64, seqNr uint64, executorId models.UserAccountId) ExpGottenByLoginBonus {
	id := uuid.Must(uuid.NewV7())
	now := time.Now()
	occurredAt := uint64(now.UnixNano() / 1e6)
	return ExpGottenByLoginBonus{
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

func NewExpGottenByLoginBonusFrom(id string, aggregateID models.PymonID, pymonID string, exp int64, seqNr uint64, executorId models.UserAccountId, occurredAt uint64) ExpGottenByLoginBonus {
	return ExpGottenByLoginBonus{
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

func (e *ExpGottenByLoginBonus) ToProto() []byte {

	p := &protobuf.EventEnvelope{
		Metadata: &protobuf.EventMetadata{
			Id:          e.id,
			AggregateId: e.aggregateID.GetValue(),
			SeqNr:       e.seqNr,
			OccurredAt:  e.occurredAt,
			ExecutorId:  e.executorId.GetValue(),
		},
		Event: &protobuf.EventEnvelope_ExpGottenByLoginBonus{
			ExpGottenByLoginBonus: &pymon.ExpGottenByLoginBonus{
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

func (e *ExpGottenByLoginBonus) GetId() string {
	return e.id
}

func (e *ExpGottenByLoginBonus) GetTypeName() string {
	return "ExpGottenByLoginBonus"
}

func (e *ExpGottenByLoginBonus) GetAggregateId() esa.AggregateId {
	return &e.aggregateID
}

func (e *ExpGottenByLoginBonus) GetSeqNr() uint64 {
	return e.seqNr
}

func (e *ExpGottenByLoginBonus) GetExecutorId() *models.UserAccountId {
	return &e.executorId
}

func (e *ExpGottenByLoginBonus) IsCreated() bool {
	return true
}

func (e *ExpGottenByLoginBonus) GetOccurredAt() uint64 {
	return e.occurredAt
}

func (g *ExpGottenByLoginBonus) String() string {
	return fmt.Sprintf("%s{ id: %s, aggregateId: %s, seqNr: %d, occurredAt: %d}",
		g.GetTypeName(), g.id, g.aggregateID, g.seqNr, g.occurredAt)
}
