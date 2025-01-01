package repository

import (
	"context"
	"database/sql"

	"github.com/K-Kizuku/pymon-graphql/internal/db/model"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/entity"
	"github.com/K-Kizuku/pymon-graphql/internal/domain/repository"
	"github.com/K-Kizuku/pymon-graphql/pkg/db"
)

type UserRepository struct {
	db *db.DB
}

var _ repository.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *db.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	user := &model.User{}
	err := u.db.DB.GetContext(ctx, user, "SELECT * FROM users WHERE id = ?", userID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}
