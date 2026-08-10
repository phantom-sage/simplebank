package gapi

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	db "github.com/phantom-sage/simplebank/db/sqlc"
	"github.com/phantom-sage/simplebank/pb"
	"github.com/phantom-sage/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPasswrod(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	user, err := server.store.CreateUser(ctx, db.CreateUserParams{
		Username:       req.GetUsername(),
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
		HashedPassword: hashedPassword,
	})
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", pgErr)
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}
	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return resp, nil
}
