
test: 
	go mod tidy
	- make docker/up
	go clean -cache
	go test -failfast $$(go list ./...) -cover
	- make docker/down

run:
	go run src/main.go

docker/up:
	docker-compose up -d

docker/down:
	docker-compose down