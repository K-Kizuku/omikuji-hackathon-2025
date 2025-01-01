package repository

import (
	"context"

	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
)

type IPythonStatRepository interface {
	GetPythonStatByID(ctx context.Context, pythonID string) (*entity.PythonStat, error)
}
