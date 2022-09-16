test:
	go test -v -cover -coverprofile=coverage/coverage.out ./...

server:
	swag init
	go run main.go

.PHONY.: test server