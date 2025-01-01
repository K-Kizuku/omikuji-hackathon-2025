ENV_FILE := .env
include $(ENV_FILE)

export $(shell sed 's/=.*//' $(ENV_FILE))

MYSQL_USERNAME := $(MYSQL_USERNAME)
MYSQL_PASSWORD := $(MYSQL_PASSWORD)
MYSQL_DATABASE := $(MYSQL_DATABASE)
MYSQL_PORT := $(MYSQL_PORT)
MYSQL_HOST := $(MYSQL_HOST)
MYSQL_TLS := $(MYSQL_TLS)

gen-migrate:
	migrate create -ext sql -dir db/mysql/migrations -seq $(name)

migrate:
	migrate -path db/mysql/migrations -database "mysql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?tls=true&multiStatements=true" up

local-migrate:
	migrate -path db/mysql/migrations -database "mysql://user:password@tcp(localhost:3306)/pymon?tls=false&multiStatements=true" up

migrate-down:
	migrate -path db/mysql/migrations -database "mysql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)?tls=true&multiStatements=true" down