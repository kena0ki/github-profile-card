.PHONY: dev, build, up, down, test, cover

dev:
	docker-compose run --rm --service-ports api bash

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

test:
	go test ./...

cover:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out
