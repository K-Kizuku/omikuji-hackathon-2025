package events

import (
	"fmt"
	"time"

	"github.com/K-Kizuku/pymon-command-api/internal/domain/models"
	"github.com/K-Kizuku/pymon-command-api/pkg/null"
	"github.com/K-Kizuku/pymon-command-api/pkg/protobuf"
	"github.com/K-Kizuku/pymon-command-api/pkg/protobuf/pymon"
	"github.com/google/uuid"
	esa "github.com/j5ik2o/event-store-adapter-go/pkg"
	"google.golang.org/protobuf/proto"
)

type PymonCreated struct {
	EventMetadata
	aggregateID  models.PymonID
	pymonID      string
	ownerID      models.UserAccountId
	RepositoryID string
	name         string
	exp          int64
	Status       models.Status
	skill1       null.Null[int32]
	skill2       null.Null[int32]
	skill3       null.Null[int32]
	skill4       null.Null[int32]
}

func NewPymonCreated(aggregateID models.PymonID, pymonID string, ownerID models.UserAccountId, RepositoryID string, name string, Status models.Status, seqNr uint64, executorId models.UserAccountId) PymonCreated {
	id := uuid.Must(uuid.NewV7())
	now := time.Now()
	occurredAt := uint64(now.UnixNano() / 1e6)
	return PymonCreated{
		EventMetadata: EventMetadata{
			id:         id.String(),
			seqNr:      seqNr,
			executorId: executorId,
			occurredAt: occurredAt,
		},
		aggregateID:  aggregateID,
		pymonID:      pymonID,
		ownerID:      ownerID,
		RepositoryID: RepositoryID,
		name:         name,
		exp:          0,
		Status:       Status,
		skill1:       null.New[int32](nil),
		skill2:       null.New[int32](nil),
		skill3:       null.New[int32](nil),
		skill4:       null.New[int32](nil),
	}
}

func NewPymonCreatedFrom(id string, aggregateID models.PymonID, pymonID string, ownerID models.UserAccountId, RepositoryID string, name string, exp int64, Status models.Status, skill1 null.Null[int32], skill2 null.Null[int32], skill3 null.Null[int32], skill4 null.Null[int32], seqNr uint64, executorId models.UserAccountId, occurredAt uint64) PymonCreated {
	return PymonCreated{
		EventMetadata: EventMetadata{
			id:         id,
			seqNr:      seqNr,
			executorId: executorId,
			occurredAt: occurredAt,
		},
		aggregateID:  aggregateID,
		pymonID:      pymonID,
		ownerID:      ownerID,
		RepositoryID: RepositoryID,
		name:         name,
		exp:          exp,
		Status:       Status,
		skill1:       skill1,
		skill2:       skill2,
		skill3:       skill3,
		skill4:       skill4,
	}
}

func (e *PymonCreated) ToProto() []byte {
	p := &protobuf.EventEnvelope{
		Metadata: &protobuf.EventMetadata{
			Id:          e.id,
			AggregateId: e.aggregateID.GetValue(),
			SeqNr:       e.seqNr,
			OccurredAt:  e.occurredAt,
			ExecutorId:  e.executorId.GetValue(),
		},
		Event: &protobuf.EventEnvelope_PymonCreated{
			PymonCreated: &pymon.PymonCreated{
				PymonId:      e.pymonID,
				OwnerId:      e.ownerID.GetValue(),
				RepositoryId: e.RepositoryID,
				Name:         e.name,
				Exp:          e.exp,
				Status: &pymon.Status{
					Hp:      e.Status.HP,
					Attack:  e.Status.Attack,
					Defense: e.Status.Defense,
					Speed:   e.Status.Speed,
				},
				Skill1: e.skill1.GetValueOptional(),
				Skill2: e.skill2.GetValueOptional(),
				Skill3: e.skill3.GetValueOptional(),
				Skill4: e.skill4.GetValueOptional(),
			},
		},
	}
	b, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	return b
}

func (e *PymonCreated) GetId() string {
	return e.id
}

func (e *PymonCreated) GetTypeName() string {
	return "PymonCreated"
}

func (e *PymonCreated) GetAggregateId() esa.AggregateId {
	return &e.aggregateID
}

func (e *PymonCreated) GetSeqNr() uint64 {
	return e.seqNr
}

func (e *PymonCreated) GetExecutorId() *models.UserAccountId {
	return &e.executorId
}

func (e *PymonCreated) IsCreated() bool {
	return true
}

func (e *PymonCreated) GetOccurredAt() uint64 {
	return e.occurredAt
}

func (e *PymonCreated) String() string {
	return fmt.Sprintf("%s{ id: %s, aggregateId: %s, seqNr: %d, occurredAt: %d}",
		e.GetTypeName(), e.id, e.aggregateID, e.seqNr, e.occurredAt)
}
