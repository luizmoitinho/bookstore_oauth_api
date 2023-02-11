
test: 
	- make docker/up
	go test -failfast $$(go list ./...) -cover

run:
	go run src/main.go

docker/up:
	docker-compose up -d

docker/up:
	docker-compose down