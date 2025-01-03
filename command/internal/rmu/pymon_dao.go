package rmu

import (
	"context"
	"time"

	"github.com/K-Kizuku/pymon-command-api/internal/domain/models"
)

type PymonDao interface {
	InsertPymon(ctx context.Context, aggregateID string, pymonID string, name string, OwnerID string, repositoryID string, exp int64, createdAt time.Time) error
	UpdatePymon(ctx context.Context, pymonID string, name string, updatedAt time.Time) error
	DeletePymon(ctx context.Context, pymonID string, deletedAt time.Time) error

	UpdateExp(ctx context.Context, pymonID string, exp int64, updatedAt time.Time) error

	UpdateSkill(ctx context.Context, pymonID string, skillID int64, skillIndex int32, updatedAt time.Time) error
	DeleteSkill(ctx context.Context, pymonID string, skillIndex int32, updatedAt time.Time) error

	UpdateStatus(ctx context.Context, pymonID string, status models.Status, updatedAt time.Time) error

	InsertUser(ctx context.Context, aggregateID string, userID string, name, githubID string, createdAt time.Time) error
	UpdateUserName(ctx context.Context, userID string, name string, updatedAt time.Time) error
	UpdateUserGithubID(ctx context.Context, userID string, githubID string, updatedAt time.Time) error
	DeleteUser(ctx context.Context, userID string, deletedAt time.Time) error
}
