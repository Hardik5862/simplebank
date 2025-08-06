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

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Hardik5862/simplebank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/Hardik5862/simplebank/worker TaskDistributor


proto:
	rm -rf pb/*.go
	rm -rf doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=./doc/swagger --openapiv2_opt=allow_merge=true \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

.PHONY: createdb dropdb migrateup migratedown migrateuplatest migratedownlatest new_migration sqlc test server mock proto
