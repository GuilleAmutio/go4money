createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres go4money

dropdb:
	docker exec -it postgres dropdb go4money

migrateup:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/go4money?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/go4money?sslmode=disable" -verbose down

sqlc:
	sqlc generate

dbtest:
	go test -v -cover ./db/sqlc

.PHONY.: postgres createdb dropdb migrateup migratedown sqlc dbtest