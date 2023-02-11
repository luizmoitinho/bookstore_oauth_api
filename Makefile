
test: 
	- make docker/up
	@go test -failfast $$(go list ./...) -cover

run:
	go run src/main.go

docker/up:
	docker compose -f docker-compose.yml up -d

docker/down:
	docker-compose down