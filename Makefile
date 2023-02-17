
test: 
	- make docker/up
	@go clean -cache
	@go test -failfast $$(go list ./...) -cover

run:
	go run src/main.go

docker/up:
	docker-compose up -d

docker/down:
	docker-compose down