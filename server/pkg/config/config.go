package config

import (
	"cmp"
	"fmt"
	"os"
)

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
	return &DBConfig{
		Host:     cmp.Or(os.Getenv("MYSQL_HOST"), "db"),
		Port:     cmp.Or(os.Getenv("MYSQL_PORT"), "3306"),
		User:     cmp.Or(os.Getenv("MYSQL_USERNAME"), "user"),
		Password: cmp.Or(os.Getenv("MYSQL_PASSWORD"), "password"),
		DBName:   cmp.Or(os.Getenv("MYSQL_DATABASE"), "pymon"),
		UseTLS:   cmp.Or(os.Getenv("MYSQL_TLS"), "false"),
	}
}
