include .env
export

service-run:
	go run main.go

migrate-up:
	migrate -path migrations -database ${CONN_STRING} up

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down

migrate-version:
	migrate -path migrations -database ${CONN_STRING} version

docker-build-go:
	docker build . -t links-service

docker-run-go:
	docker run -d -p 5000:5000 links-service

//FIXME
docker-run-postgres-v: 
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} --volume ./out/pgdata:/var/lib/postgresql/data postgres:17-bookworm 

docker-run-postgres:
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} postgres:18-bookworm 