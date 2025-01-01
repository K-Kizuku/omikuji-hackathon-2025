package entity

import "time"

type User struct {
	ID        string
	Name      string
	GithubID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDateData struct {
	LoginedAt time.Time
}

func (u *UserDateData) IsLoginedToday(now time.Time) bool {
	return u.LoginedAt.Year() == now.Year() && u.LoginedAt.YearDay() == now.YearDay()
}

func (u *UserDateData) Login(now time.Time) {
	u.LoginedAt = now
}
