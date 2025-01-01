package model

import (
	"time"

	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
)

type Python struct {
	ID         string    `db:"id"`
	UserID     string    `db:"user_id"`
	Name       string    `db:"name"`
	Repository string    `db:"repository"`
	Exp        int       `db:"exp"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func NewPythonModelFromEntity(python *entity.Python) *Python {
	return &Python{
		ID:         python.ID,
		UserID:     python.UserID,
		Name:       python.Name,
		Repository: python.Repository,
		Exp:        python.Exp,
		CreatedAt:  python.CreatedAt,
		UpdatedAt:  python.UpdatedAt,
	}
}

func (p *Python) ToEntity() *entity.Python {
	return &entity.Python{
		ID:         p.ID,
		UserID:     p.UserID,
		Name:       p.Name,
		Repository: p.Repository,
		Exp:        p.Exp,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}
