# Name of your Docker Compose service
SERVICE_NAME=discordbot

# Default log file path
LOG_FILE=logs/discordbot.json

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

# View the logs in real time
logs:
	tail -f $(LOG_FILE)

# View the last 50 lines of logs
logs-last:
	tail -n 50 $(LOG_FILE)

# Clear the log file
logs-clean:
	@echo "[]" > $(LOG_FILE)
	@echo "✅ Log file cleaned."

# Open a shell in the container
shell:
	docker exec -it $(SERVICE_NAME) sh
