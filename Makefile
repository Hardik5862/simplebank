createdb:
	docker compose exec -it db createdb --username postgres --owner postgres simplebank

dropdb:
	docker compose exec -it db dropdb --username postgres simplebank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Hardik5862/simplebank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc test server mock
