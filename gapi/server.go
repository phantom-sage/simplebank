package gapi

import (
	"fmt"

	db "github.com/phantom-sage/simplebank/db/sqlc"
	"github.com/phantom-sage/simplebank/pb"
	"github.com/phantom-sage/simplebank/token"
	"github.com/phantom-sage/simplebank/util"
	"github.com/phantom-sage/simplebank/worker"
)

// Server serve gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer create new server instance.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKEy)
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
