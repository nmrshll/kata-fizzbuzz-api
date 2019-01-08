dev: restart-api logs-api

restart-api:
	docker-compose kill api && docker-compose rm -v -f api && docker-compose up -d --no-deps --build api

build-api:
	docker-compose build api

ps:
	docker-compose ps
up:
	docker-compose up -d
down:
	docker-compose down

logs:
	docker-compose logs -f --tail=100
logs-api:
	docker-compose logs -f --tail=100 api

test:
	go test ./...