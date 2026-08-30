include .env
export

docker-run-postgres:
	docker compose up postgres

docker-stop-postgres:
	docker compose down postgres

docker-run-go:
	docker compose up backend

docker-stop-go:
	docker compose down backend

migrate-up:
	migrate -path migrations -database ${CONN_STRING_MIGRATE} up

migrate-down:
	migrate -path migrations -database ${CONN_STRING_MIGRATE} down

migrate-version:
	migrate -path migrations -database ${CONN_STRING_MIGRATE} version

docker-run:
	make docker-run-postgres && make migrate-up && make docker-run-go

docker-stop:
	docker compose down


