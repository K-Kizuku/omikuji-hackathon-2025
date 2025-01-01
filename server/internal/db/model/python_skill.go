package model

import "github.com/K-Kizuku/pymon-graphql/internal/domain/entity"

type PythonSkill struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Pp          int    `db:"pp"`
	Attack      int    `db:"attack"`
	HitRate     int    `db:"hit_rate"`
	MinVersion  string `db:"min_version"`
	MaxVersion  string `db:"max_version"`
}

func NewPythonSkillModelFromEntity(pythonSkill *entity.PythonSkill) *PythonSkill {
	return &PythonSkill{
		ID:          pythonSkill.ID,
		Name:        pythonSkill.Name,
		Description: pythonSkill.Description,
		Pp:          pythonSkill.Pp,
		Attack:      pythonSkill.Attack,
		HitRate:     pythonSkill.HitRate,
		MinVersion:  pythonSkill.MinVersion,
		MaxVersion:  pythonSkill.MaxVersion,
	}
}

func (p *PythonSkill) ToEntity() *entity.PythonSkill {
	return &entity.PythonSkill{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Pp:          p.Pp,
		Attack:      p.Attack,
		HitRate:     p.HitRate,
		MinVersion:  p.MinVersion,
		MaxVersion:  p.MaxVersion,
	}
}
