package gapi

import (
	"fmt"

	db "github.com/Hardik5862/simplebank/db/sqlc"
	"github.com/Hardik5862/simplebank/pb"
	"github.com/Hardik5862/simplebank/token"
	"github.com/Hardik5862/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimplebankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
