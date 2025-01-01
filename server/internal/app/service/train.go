package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/service"
	"github.com/redis/go-redis/v9"
)

type TrainService struct {
	redis *redis.Client
}

var _ service.ITrainService = (*TrainService)(nil)

func NewTrainService(redis *redis.Client) *TrainService {
	return &TrainService{redis: redis}
}

// TODO: 日付が変わる毎にRemainTrainCountをリセットする処理を追加

func (t *TrainService) RemainTrainCount(ctx context.Context, pythonID string) (int, error) {
	data, err := t.redis.Get(ctx, pythonID).Bytes()
	if err != nil && err != redis.Nil {
		return -1, err
	}
	var train entity.Train
	if err == redis.Nil {
		return entity.MaxTrainCount, nil
	}
	if err := json.Unmarshal(data, &train); err != nil {
		return -1, err
	}
	return train.RemainTrainCount, nil
}

func (t *TrainService) Train(ctx context.Context, pythonID string) error {
	data, err := t.redis.Get(ctx, pythonID).Bytes()
	if err != nil && err != redis.Nil {
		return err
	}
	var train entity.Train
	if err == redis.Nil {
		train.ResetCount()
	} else {
		if err := json.Unmarshal(data, &train); err != nil {
			return err
		}
	}
	if !train.CanTrain() {
		return errors.New("train count is not enough")
	}
	train.Train()
	data, err = json.Marshal(train)
	if err != nil {
		return err
	}
	// TODO: PythonのExpを増やす処理を追加
	return t.redis.Set(ctx, pythonID, data, 24*time.Hour).Err()
}
