package db

import "context"

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return CreateUserTxResult{}, err
	}

	defer tx.Rollback(ctx)
	qtx := store.WithTx(tx)

	var result CreateUserTxResult
	result.User, err = qtx.CreateUser(ctx, arg.CreateUserParams)
	if err != nil {
		return CreateUserTxResult{}, err
	}
	err = arg.AfterCreate(result.User)
	if err != nil {
		return CreateUserTxResult{}, err
	}

	return result, tx.Commit(ctx)
}
