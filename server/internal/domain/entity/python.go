package entity

import "time"

type Python struct {
	ID         string
	Name       string
	Exp        int
	Repository string
	UserID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
