package repository

import (
	"context"

	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
)

type IPythonSkillRepository interface {
	SelectByIDs(ctx context.Context, ids []int) ([]*entity.PythonSkill, error)
}
