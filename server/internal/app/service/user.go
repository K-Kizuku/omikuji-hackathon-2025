package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/service"
	"github.com/redis/go-redis/v9"
)

type userIDKey struct{}

type UserService struct {
	redis *redis.Client
}

var _ service.IUserService = (*UserService)(nil)

func NewUserService(redis *redis.Client) *UserService {
	return &UserService{redis: redis}
}

func (u *UserService) IsLogined(ctx context.Context, userID string) (bool, error) {
	now := time.Now()
	data, err := u.redis.Get(ctx, userID).Bytes()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	var userData entity.UserDateData
	if err := json.Unmarshal(data, &userData); err != nil {
		return false, err
	}
	return userData.IsLoginedToday(now), nil
}

func (u *UserService) Login(ctx context.Context, userID string) error {
	now := time.Now()
	var userData entity.UserDateData
	userData.Login(now)
	data, err := json.Marshal(userData)
	if err != nil {
		return err
	}
	return u.redis.Set(ctx, userID, data, 24*time.Hour).Err()
}
