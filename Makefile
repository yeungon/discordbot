.PHONY: up down run dev update pull

up:
	docker-compose up --build
down:
	docker-compose down
run:
	docker-compose up -d --build
dev:
	go run ./cmd
pull:
	sudo git pull

update: pull run
