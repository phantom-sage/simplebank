package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type VerifyEmailTxParams struct {
	EmailId    int64
	SecretCode string
}

type VerifyEmailTxResult struct {
	User        User
	VerifyEmail VerifyEmail
}

func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return VerifyEmailTxResult{}, err
	}

	defer tx.Rollback(ctx)
	qtx := store.WithTx(tx)

	var result VerifyEmailTxResult
	result.VerifyEmail, err = qtx.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
		ID:         arg.EmailId,
		SecretCode: arg.SecretCode,
	})
	if err != nil {
		return VerifyEmailTxResult{}, err
	}

	result.User, err = qtx.UpdateUser(ctx, UpdateUserParams{
		Username: result.VerifyEmail.Username,
		IsEmailVerified: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
	})
	if err != nil {
		return VerifyEmailTxResult{}, err
	}

	return result, tx.Commit(ctx)
}
