DB_URL=postgresql://postgres:postgres@localhost:5432/simplebank?sslmode=disable

createdb:
	docker compose exec -it db createdb --username postgres --owner postgres simplebank

dropdb:
	docker compose exec -it db dropdb --username postgres simplebank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateuplatest:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedownlatest:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Hardik5862/simplebank/db/sqlc Store

proto:
	rm -rf pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: createdb dropdb migrateup migratedown migrateuplatest migratedownlatest sqlc test server mock proto
