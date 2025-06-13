.PHONY: up down run dev

up:
	docker-compose up --build
down:
	docker-compose down
run:
	docker-compose up -d --build
dev:
	go run ./cmd
