package rmu

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/K-Kizuku/pymon-command-api/pkg/protobuf"

	dynamodbevents "github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

type ReadModelUpdater struct {
	dao PymonDao
}

func NewReadModelUpdater(dao PymonDao) ReadModelUpdater {
	return ReadModelUpdater{dao}
}

func (r *ReadModelUpdater) UpdateReadModel(ctx context.Context, event dynamodbevents.DynamoDBEvent) (string, error) {
	res := "OK"
	for _, record := range event.Records {
		slog.Info("Processing request data for event GetId %s, type %s.", record.EventID, record.EventName)
		attributeValues := record.Change.NewImage
		payloadBytes := convertToBytes(attributeValues["payload"])
		event, err := convertToEventEnvelope(payloadBytes)
		if err != nil {
			return res, err
		}
		slog.Debug(fmt.Sprintf("EventEnvelope = %+v", event))
		switch ev := event.Event.(type) {
		case *protobuf.EventEnvelope_ExpGottenByLoginBonus:
			if err := addExpByLoginBonus(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_ExpGottenByTraining:
			if err := addExpByTraining(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_PymonCreated:
			if err := createPymon(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_PymonNameChanged:
			if err := changePymonName(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_PymonDeleted:
			if err := deletePymon(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_SpAssigned:
			if err := updatePymonStatus(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_UserCreated:
			if err := createUser(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_UserNameChanged:
			if err := changeUserName(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_UserDeleted:
			if err := deleteUser(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_SkillLearned:
			if err := learnSkill(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		case *protobuf.EventEnvelope_SkillForgotten:
			if err := forgetSkill(ctx, event.Metadata, ev, r); err != nil {
				return res, err
			}
		default:
			slog.Error(fmt.Sprintf("Unknown event type: %T", ev))
		}
		for name, value := range record.Change.NewImage {
			slog.Debug(fmt.Sprintf("Attribute name: %s, value: %s", name, value.String()))
		}
	}
	return res, nil
}
