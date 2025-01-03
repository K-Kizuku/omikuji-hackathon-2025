package main

import (
	"context"
	"fmt"

	"github.com/K-Kizuku/pymon-command-api/pkg/db"
	"github.com/alecthomas/kong"
	"github.com/fujiwara/lamblocal"
	"github.com/jmoiron/sqlx"
)

type CLI struct {
	DB struct {
		Host     string `help:"Database host." default:"db" env:"MYSQL_HOST" required:""`
		Port     int    `help:"Database port." default:"3306" env:"MYSQL_PORT" required:""`
		User     string `help:"Database user." default:"user" env:"MYSQL_USERNAME" required:""`
		Password string `help:"Database password." default:"password" env:"MYSQL_PASSWORD" required:""`
		DBName   string `help:"Database name." default:"pymon" env:"MYSQL_DATABASE" required:""`
		TLS      bool   `help:"Use TLS." default:"false" env:"MYSQL_TLS" required:""`
	}
}

func hello(_ context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello %s", name), nil
}

func main() {
	ctx := context.Background()
	var c CLI
	kong.Parse(&c)
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&tls=%t", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DBName, c.DB.TLS)
	db := db.NewDB(dbDSN)
	defer func(db *sqlx.DB) {
		if db != nil {
			err := db.Close()
			if err != nil {
				panic(err.Error())
			}
		}
	}(db.DB)
	lamblocal.Run(ctx, hello)
}
