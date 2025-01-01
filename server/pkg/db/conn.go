package db

import (
	"log"

	"github.com/K-Kizuku/pymon-graphql/pkg/config"
	"github.com/jmoiron/sqlx"
)

const driverName = "mysql"

type DB struct {
	*sqlx.DB
}

func NewDB(cfg *config.DBConfig) *DB {
	db, err := sqlx.Open(driverName, cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &DB{db}
}
