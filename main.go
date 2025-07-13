package main

import (
	"context"
	"log"

	"github.com/Hardik5862/simplebank/api"
	db "github.com/Hardik5862/simplebank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://postgres:postgres@localhost:5432/simplebank?sslmode=disable"
	serverAddress = "0.0.0.0:8005"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
