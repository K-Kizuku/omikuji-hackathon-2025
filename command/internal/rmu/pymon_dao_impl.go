package rmu

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/K-Kizuku/pymon-command-api/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type PymonDaoImpl struct {
	db *sqlx.DB
}

func NewPymonDaoImpl(db *sqlx.DB) PymonDao {
	return PymonDaoImpl{db}
}

// User
func (p PymonDaoImpl) InsertUser(ctx context.Context, aggregateID string, userID string, name, githubID string, createdAt time.Time) error {
	stmt, err := p.db.PrepareContext(ctx, "INSERT INTO users (id, name, githubID, created_at, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	dt := createdAt.Format("2006-01-02 15:04:05")
	_, err = stmt.ExecContext(ctx, aggregateID, false, name, githubID, dt, dt)
	if err != nil {
		return err
	}
	return nil
}

func (p PymonDaoImpl) UpdateUserName(ctx context.Context, userID string, name string, updatedAt time.Time) error {
	stmt, err := p.db.PrepareContext(ctx, "UPDATE users SET name = ?, updated_at = ?  WHERE id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	dt := updatedAt.Format("2006-01-02 15:04:05")
	_, err = stmt.ExecContext(ctx, name, dt, userID)
	if err != nil {
		return err
	}
	return nil
}

func (p PymonDaoImpl) UpdateUserGithubID(ctx context.Context, userID string, githubID string, updatedAt time.Time) error {
	stmt, err := p.db.PrepareContext(ctx, "UPDATE users SET github_id = ?, updated_at = ?  WHERE id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	dt := updatedAt.Format("2006-01-02 15:04:05")
	_, err = stmt.ExecContext(ctx, githubID, dt, userID)
	if err != nil {
		return err
	}
	return nil
}

func (p PymonDaoImpl) DeleteUser(ctx context.Context, userID string, deletedAt time.Time) error {
	panic("unimplemented")
}

// Pymon
func (p PymonDaoImpl) InsertPymon(ctx context.Context, aggregateID string, pymonID string, name string, OwnerID string, repositoryID string, exp int64, createdAt time.Time) error {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()
		if err != nil {
			slog.Error(err.Error())
		}
	}(tx)

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO pythons (id, user_id, name, repository, exp, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	dt := createdAt.Format("2006-01-02 15:04:05")
	_, err = stmt.ExecContext(ctx, pymonID, OwnerID, name, repositoryID, exp, dt, dt)
	if err != nil {
		return err
	}

	stmt, err = tx.PrepareContext(ctx, "INSERT INTO pythons_stats (id) VALUES (?)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	_, err = stmt.ExecContext(ctx, pymonID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return nil
}

func (p PymonDaoImpl) UpdatePymon(ctx context.Context, pymonID string, name string, updatedAt time.Time) error {
	stmt, err := p.db.PrepareContext(ctx, "UPDATE pythons SET name = ?, updated_at = ?  WHERE id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	dt := updatedAt.Format("2006-01-02 15:04:05")
	_, err = stmt.ExecContext(ctx, name, dt, pymonID)
	if err != nil {
		return err
	}
	return nil
}

func (p PymonDaoImpl) DeletePymon(ctx context.Context, pymonID string, deletedAt time.Time) error {
	panic("unimplemented")
}

// Skill
func (p PymonDaoImpl) UpdateSkill(ctx context.Context, pymonID string, skillID int64, skillIndex int32, updatedAt time.Time) error {
	stmt, err := p.db.PrepareContext(ctx, fmt.Sprintf("UPDATE pythons_stats SET skill%d = ?  WHERE id = ?", skillIndex))
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	_, err = stmt.ExecContext(ctx, skillID, pymonID)
	if err != nil {
		return err
	}
	return nil
}

func (p PymonDaoImpl) DeleteSkill(ctx context.Context, pymonID string, skillIndex int32, updatedAt time.Time) error {
	stmt, err := p.db.PrepareContext(ctx, fmt.Sprintf("UPDATE pythons_stats SET skill%d = NULL WHERE id = ?", skillIndex))
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	_, err = stmt.ExecContext(ctx, pymonID)
	if err != nil {
		return err
	}
	return nil
}

// Status
func (p PymonDaoImpl) UpdateStatus(ctx context.Context, pymonID string, status models.Status, updatedAt time.Time) error {
	stmt, err := p.db.PrepareContext(ctx, "UPDATE pythons_stats SET hp = hp + ?, attack = attack + ?, defense = defence + ?, speed = speed + ?  WHERE id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	_, err = stmt.ExecContext(ctx, status.HP, status.Attack, status.Defense, status.Speed, pymonID)
	if err != nil {
		return err
	}
	return nil
}

// Exp
func (p PymonDaoImpl) UpdateExp(ctx context.Context, pymonID string, exp int64, updatedAt time.Time) error {
	stmt, err := p.db.PrepareContext(ctx, "UPDATE pythons SET exp = exp + ?, updated_at = ?  WHERE id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)
	dt := updatedAt.Format("2006-01-02 15:04:05")
	_, err = stmt.ExecContext(ctx, exp, dt, pymonID)
	if err != nil {
		return err
	}
	return nil
}
