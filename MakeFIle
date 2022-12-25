PWD = ${CURDIR}
NAME = web-2022

# Запустить проект
.PHONY: run
run:
	go run $(PWD)/cmd/$(NAME)/

# Сбилдить проект
.PHONY: build
build:
	go build -o bin/$(NAME) $(PWD)/cmd/$(NAME)

# Создать .env файл
.PHONY: local
local:
	cp .dist.env .env

# Запустить миграции
.PHONY: migrate
migrate:
	go run $(PWD)/cmd/migrate

# Запустить docker
.PHONY: docker
docker:
	docker compose up -d

# Сгенерировать swagger
.PHONY: swagger
swagger:
	swag init -g cmd/$(NAME)/main.go