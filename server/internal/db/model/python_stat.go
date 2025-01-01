package model

import (
	"database/sql"

	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
)

type PythonStat struct {
	PythonID string        `db:"python_id"`
	Hp       int           `db:"hp"`
	Attack   int           `db:"attack"`
	Defense  int           `db:"defense"`
	Speed    int           `db:"speed"`
	Skill1   sql.NullInt64 `db:"skill1"`
	Skill2   sql.NullInt64 `db:"skill2"`
	Skill3   sql.NullInt64 `db:"skill3"`
	Skill4   sql.NullInt64 `db:"skill4"`
}

func NewPythonStatsModelFromEntity(stats *entity.PythonStat) *PythonStat {
	return &PythonStat{
		PythonID: stats.PythonID,
		Hp:       stats.Hp,
		Attack:   stats.Attack,
		Defense:  stats.Defense,
		Speed:    stats.Speed,
		Skill1: sql.NullInt64{
			Int64: int64(*stats.Skill1),
			Valid: stats.Skill1 != nil,
		},
		Skill2: sql.NullInt64{
			Int64: int64(*stats.Skill2),
			Valid: stats.Skill2 != nil,
		},
		Skill3: sql.NullInt64{
			Int64: int64(*stats.Skill3),
			Valid: stats.Skill3 != nil,
		},
		Skill4: sql.NullInt64{
			Int64: int64(*stats.Skill4),
			Valid: stats.Skill4 != nil,
		},
	}
}

func sqlNullInt64ToInt(v sql.NullInt64) *int {
	if v.Valid {
		i := int(v.Int64)
		return &i
	}
	return nil
}

func (s *PythonStat) ToEntity() *entity.PythonStat {
	return &entity.PythonStat{
		PythonID: s.PythonID,
		Hp:       s.Hp,
		Attack:   s.Attack,
		Defense:  s.Defense,
		Speed:    s.Speed,
		Skill1:   sqlNullInt64ToInt(s.Skill1),
		Skill2:   sqlNullInt64ToInt(s.Skill2),
		Skill3:   sqlNullInt64ToInt(s.Skill3),
		Skill4:   sqlNullInt64ToInt(s.Skill4),
	}
}
