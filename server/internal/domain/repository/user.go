package repository

import (
	"context"

	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
)

type IUserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}
