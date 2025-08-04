package gapi

import (
	"fmt"

	db "github.com/Hardik5862/simplebank/db/sqlc"
	"github.com/Hardik5862/simplebank/pb"
	"github.com/Hardik5862/simplebank/token"
	"github.com/Hardik5862/simplebank/util"
	"github.com/Hardik5862/simplebank/worker"
)

type Server struct {
	pb.UnimplementedSimplebankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
