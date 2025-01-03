package db

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

type DB struct {
	*sqlx.DB
}

func NewDB(dbDSN string) *DB {
	db, err := sqlx.Open(driverName, dbDSN)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &DB{db}
}
