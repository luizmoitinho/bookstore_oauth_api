
test: 
	go test -failfast $$(go list ./...) -cover

run:
	go run src/main.go