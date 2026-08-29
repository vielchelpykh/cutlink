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

docker-build:
	docker build . -t links-service

docker-run:
	docker run -d -p 5000:5000 links-service
