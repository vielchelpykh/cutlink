include .env
export

service-run:
	go run main.go

migrate-up:
	migrate -path migrations -database ${CONN_STRING_MIGRATE} up

migrate-down:
	migrate -path migrations -database ${CONN_STRING_MIGRATE} down

migrate-version:
	migrate -path migrations -database ${CONN_STRING_MIGRATE} version

# docker-build-go:
# 	docker build . -t links-service

# docker-run-go:
# 	docker run -d -p 5000:5000 links-service

# docker-run-postgres-v: 
# 	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} --volume ./out/pgdata:/var/lib/postgresql/data postgres:alpine

# docker-run-postgres:
# 	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} postgres:18-bookworm 

docker-run-go:
	docker compose up backend

docker-stop-go:
	docker compose down backend

docker-run-postgres:
	docker compose up postgres

docker-stop-postgres:
	docker compose down postgres