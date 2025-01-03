package rmu

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/K-Kizuku/pymon-command-api/internal/domain/models"
	"github.com/K-Kizuku/pymon-command-api/pkg/protobuf"
	dynamodbevents "github.com/aws/aws-lambda-go/events"
	"google.golang.org/protobuf/proto"
)

func addExpByLoginBonus(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_ExpGottenByLoginBonus, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("addExpByLoginBonus: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.UpdateExp(ctx, ev.ExpGottenByLoginBonus.GetPymonId(), ev.ExpGottenByLoginBonus.GetExp(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("addExpByLoginBonust: finished")
	return nil
}

func addExpByTraining(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_ExpGottenByTraining, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("addExpByTraining: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.UpdateExp(ctx, ev.ExpGottenByTraining.GetPymonId(), ev.ExpGottenByTraining.GetExp(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("addExpByTraining: finished")
	return nil
}

func createPymon(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_PymonCreated, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("createPymon: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.InsertPymon(ctx, meta.GetAggregateId(), ev.PymonCreated.GetPymonId(), ev.PymonCreated.GetName(), ev.PymonCreated.GetOwnerId(), ev.PymonCreated.GetRepositoryId(), ev.PymonCreated.GetExp(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("createPymon: finished")
	return nil
}

func changePymonName(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_PymonNameChanged, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("changePymonName: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.UpdatePymon(ctx, ev.PymonNameChanged.GetPymonId(), ev.PymonNameChanged.GetName(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("changePymonName: finished")
	return nil
}

func deletePymon(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_PymonDeleted, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("deletePymon: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.DeletePymon(ctx, ev.PymonDeleted.GetId(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("deletePymon: finished")
	return nil
}

func updatePymonStatus(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_SpAssigned, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("updatePymonStatus: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	status := models.Status{
		HP:      ev.SpAssigned.GetStatus().GetHp(),
		Attack:  ev.SpAssigned.GetStatus().GetAttack(),
		Defense: ev.SpAssigned.GetStatus().GetDefense(),
		Speed:   ev.SpAssigned.GetStatus().GetSpeed(),
	}
	err := r.dao.UpdateStatus(ctx, ev.SpAssigned.GetPymonId(), status, occurredAt)
	if err != nil {
		return err
	}
	slog.Info("updatePymonStatus: finished")
	return nil
}

func createUser(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_UserCreated, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("createUser: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.InsertUser(ctx, meta.GetAggregateId(), ev.UserCreated.GetUserId(), ev.UserCreated.GetName(), ev.UserCreated.GetGithubId(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("createUser: finished")
	return nil
}

func changeUserName(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_UserNameChanged, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("changeUserName: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.UpdateUserName(ctx, ev.UserNameChanged.GetUserId(), ev.UserNameChanged.GetName(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("changeUserName: finished")
	return nil
}

func deleteUser(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_UserDeleted, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("deleteUser: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.DeleteUser(ctx, ev.UserDeleted.GetUserId(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("deleteUser: finished")
	return nil
}

func learnSkill(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_SkillLearned, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("learnSkill: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.UpdateSkill(ctx, ev.SkillLearned.GetPymonId(), ev.SkillLearned.GetSkillId(), ev.SkillLearned.GetSkillIndex(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("learnSkill: finished")
	return nil
}

func forgetSkill(ctx context.Context, meta *protobuf.EventMetadata, ev *protobuf.EventEnvelope_SkillForgotten, r *ReadModelUpdater) error {
	slog.Info(fmt.Sprintf("forgetSkill: start: ev = %v", ev))
	occurredAt := convertToTime(meta.GetOccurredAt())
	err := r.dao.DeleteSkill(ctx, ev.SkillForgotten.GetPymonId(), ev.SkillForgotten.GetSkillIndex(), occurredAt)
	if err != nil {
		return err
	}
	slog.Info("forgetSkill: finished")
	return nil
}

func convertToTime(epoc uint64) time.Time {
	occurredAtUnix := int64(epoc) * int64(time.Millisecond)
	occurredAt := time.Unix(0, occurredAtUnix)
	return occurredAt
}

func convertToBytes(payloadAttr dynamodbevents.DynamoDBAttributeValue) []byte {
	var payloadBytes []byte
	if payloadAttr.DataType() == dynamodbevents.DataTypeBinary {
		payloadBytes = payloadAttr.Binary()
	} else if payloadAttr.DataType() == dynamodbevents.DataTypeString {
		payloadBytes = []byte(payloadAttr.String())
	}
	return payloadBytes
}

func convertToEventEnvelope(bytes []byte) (*protobuf.EventEnvelope, error) {
	var p protobuf.EventEnvelope
	err := proto.Unmarshal(bytes, &p)
	if err != nil {
		slog.Info(fmt.Sprintf("getTypeString: err = %v, %s", err, string(bytes)))
		return nil, err
	}
	return &p, nil
}
