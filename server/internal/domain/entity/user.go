package entity

import "time"

type User struct {
	ID        string
	Name      string
	GithubID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
