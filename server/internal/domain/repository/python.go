package repository

import (
	"context"

	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
)

type IPythonRepository interface {
	GetPythonByUserID(ctx context.Context, userID string) (*entity.Python, error)
}
