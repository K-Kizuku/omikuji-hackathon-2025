package repository

import (
	"context"
	"database/sql"

	"github.com/K-Kizuku/pymon-graphql/internal/db/model"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/repository"
	"github.com/K-Kizuku/pymon-graphql/pkg/db"
)

type PythonRepository struct {
	db *db.DB
}

var _ repository.IPythonRepository = (*PythonRepository)(nil)

func NewPythonRepository(db *db.DB) *PythonRepository {
	return &PythonRepository{db: db}
}

func (p *PythonRepository) GetPythonByUserID(ctx context.Context, userID string) (*entity.Python, error) {
	python := &model.Python{}
	err := p.db.DB.GetContext(ctx, python, "SELECT * FROM pythons WHERE user_id = ?", userID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return python.ToEntity(), nil
}
