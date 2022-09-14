test:
	go test -v -cover -coverprofile=coverage/coverage.out ./...

server:
	go run main.go

.PHONY.: test server