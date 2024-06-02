hello:
	go test ./...

run:
	go mod vendor
	docker-compose down && docker-compose up --build
