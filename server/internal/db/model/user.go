package model

import (
	"time"

	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	GithubID  string    `db:"github_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUserModelFromEntity(user *entity.User) *User {
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		GithubID:  user.GithubID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		GithubID:  u.GithubID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
