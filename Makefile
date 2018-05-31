.PHONY: up
up:
	docker-compose build && docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: test
test:
	docker-compose build && docker-compose run shorturl go test ./... && docker-compose down
