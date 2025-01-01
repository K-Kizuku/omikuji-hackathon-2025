package firebase

import (
	"context"
	"encoding/base64"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/K-Kizuku/pymon-graphql/pkg/config"
	"google.golang.org/api/option"
)

func NewAuthApp(cfg *config.AppConfig) (*auth.Client, error) {
	decoded, err := base64.StdEncoding.DecodeString(cfg.FirebaseSecret)
	if err != nil {
		return nil, err
	}
	opt := option.WithCredentialsJSON(decoded)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
