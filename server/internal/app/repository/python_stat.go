package repository

import (
	"context"

	"github.com/K-Kizuku/pymon-graphql/internal/db/model"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/repository"
	"github.com/K-Kizuku/pymon-graphql/pkg/db"
)

type PythonStatRepository struct {
	db *db.DB
}

var _ repository.IPythonStatRepository = (*PythonStatRepository)(nil)

func NewPythonStatRepository(db *db.DB) *PythonStatRepository {
	return &PythonStatRepository{db: db}
}

func (p *PythonStatRepository) GetPythonStatByID(ctx context.Context, pythonID string) (*entity.PythonStat, error) {
	pythonStat := &model.PythonStat{}
	err := p.db.DB.GetContext(ctx, pythonStat, "SELECT * FROM python_stat WHERE python_id = ?", pythonID)
	if err != nil {
		return nil, err
	}
	return pythonStat.ToEntity(), nil
}
