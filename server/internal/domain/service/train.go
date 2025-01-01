package service

import "context"

type ITrainService interface {
	RemainTrainCount(ctx context.Context, pythonID string) (int, error)
	Train(ctx context.Context, pythonID string) error
}
