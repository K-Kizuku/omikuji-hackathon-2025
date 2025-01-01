package service

import "context"

type IUserService interface {
	IsLogined(ctx context.Context, userID string) (bool, error)
	Login(ctx context.Context, userID string) error
}
