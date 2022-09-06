createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres go4money

dropdb:
	docker exec -it postgres dropdb go4money

migrateup:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/go4money?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/go4money?sslmode=disable" -verbose down

resetdb:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/go4money?sslmode=disable" -verbose down
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/go4money?sslmode=disable" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover -coverprofile=coverage/coverage.out ./...

dbtest:
	go test -v -cover -coverprofile=coverage/dbcoverage.out ./db/sqlc

server:
	go run main.go

.PHONY.: postgres createdb dropdb migrateup migratedown resetdb sqlc test dbtest server