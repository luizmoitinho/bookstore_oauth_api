
test: 
	go mod init
	go mod tidy
	go test -failfast $$(go list ./...) -cover

run:
	go run src/main.go