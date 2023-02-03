
test: 
	go test -failfast $$(go list ./...) -cover
