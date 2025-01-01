package entity

import "time"

const (
	MaxTrainCount = 3
)

type Python struct {
	ID         string
	Name       string
	Exp        int
	Repository string
	UserID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Train struct {
	RemainTrainCount int
}

func (t *Train) CanTrain() bool {
	return t.RemainTrainCount > 0
}

func (t *Train) Train() {
	t.RemainTrainCount--
}

func (t *Train) ResetCount() {
	t.RemainTrainCount = MaxTrainCount
}
