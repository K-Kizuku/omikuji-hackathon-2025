package repository

import (
	"context"

	"github.com/K-Kizuku/pymon-graphql/internal/db/model"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/repository"
	"github.com/K-Kizuku/pymon-graphql/pkg/db"
	"github.com/jmoiron/sqlx"
)

type PythonSkillRepository struct {
	db *db.DB
}

var _ repository.IPythonSkillRepository = (*PythonSkillRepository)(nil)

func NewPythonSkillRepository(db *db.DB) *PythonSkillRepository {
	return &PythonSkillRepository{db: db}
}

func (p *PythonSkillRepository) SelectByIDs(ctx context.Context, ids []int) (entity.PythonSkills, error) {
	pythonSkills := make(model.PythonSkills, 0, len(ids))
	query := "SELECT * FROM pythons_skills WHERE id IN (?)"
	sql, params, err := sqlx.In(query, ids)
	if err != nil {
		return entity.PythonSkills{}, err
	}
	err = p.db.DB.SelectContext(ctx, &pythonSkills, sql, params...)
	if err != nil {
		return entity.PythonSkills{}, err
	}
	return pythonSkills.ToEntities(), nil
}
