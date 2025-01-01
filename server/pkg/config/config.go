package config

import (
	"cmp"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	UseTLS   string
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=%s", c.User, c.Password, c.Host, c.Port, c.DBName, c.UseTLS)
}

func NewDBConfig() *DBConfig {
	environment := cmp.Or(os.Getenv("environment"), "dev")
	switch environment {
	case "prd":
		return &DBConfig{
			Host:     cmp.Or(os.Getenv("MYSQL_HOST"), "localhost"),
			Port:     cmp.Or(os.Getenv("MYSQL_PORT"), "3306"),
			User:     cmp.Or(os.Getenv("MYSQL_USERNAME"), "user"),
			Password: cmp.Or(os.Getenv("MYSQL_PASSWORD"), "password"),
			DBName:   cmp.Or(os.Getenv("MYSQL_DATABASE"), "pymon"),
			UseTLS:   cmp.Or(os.Getenv("MYSQL_TLS"), "false"),
		}
	}
	return &DBConfig{
		Host:     cmp.Or(os.Getenv("MYSQL_DEV_HOST"), "db"),
		Port:     cmp.Or(os.Getenv("MYSQL_DEV_PORT"), "3306"),
		User:     cmp.Or(os.Getenv("MYSQL_DEV_USERNAME"), "user"),
		Password: cmp.Or(os.Getenv("MYSQL_DEV_PASSWORD"), "password"),
		DBName:   cmp.Or(os.Getenv("MYSQL_DEV_DATABASE"), "pymon"),
		UseTLS:   cmp.Or(os.Getenv("MYSQL_DEV_TLS"), "false"),
	}
}
