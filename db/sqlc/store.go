package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// TransferTxParams contains the input for transfer money transaction.
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult contains the result of transfer money transaction.
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx transfer money between accounts.
// It create transfer record, add account entries and update accounts balances within single database transaction.
func (s *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return TransferTxResult{}, err
	}
	defer tx.Rollback(ctx)
	qtx := s.WithTx(tx)

	var result TransferTxResult
	result.Transfer, err = qtx.CreateTransfer(ctx, CreateTransferParams{
		FromAccountID: pgtype.Int8{Int64: arg.FromAccountID, Valid: true},
		ToAccountID:   pgtype.Int8{Int64: arg.ToAccountID, Valid: true},
		Amount:        arg.Amount,
	})
	if err != nil {
		return TransferTxResult{}, err
	}

	result.FromEntry, err = qtx.CreateEntry(ctx, CreateEntryParams{
		AccountID: pgtype.Int8{Int64: arg.FromAccountID, Valid: true},
		Amount:    -arg.Amount,
	})
	if err != nil {
		return TransferTxResult{}, err
	}

	result.ToEntry, err = qtx.CreateEntry(ctx, CreateEntryParams{
		AccountID: pgtype.Int8{Int64: arg.ToAccountID, Valid: true},
		Amount:    arg.Amount,
	})
	if err != nil {
		return TransferTxResult{}, err
	}

	// update account with smaller ID first.
	if arg.FromAccountID < arg.ToAccountID {
		// update from account balance.
		result.FromAccount, err = qtx.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return TransferTxResult{}, err
		}

		// update to account balance.
		result.ToAccount, err = qtx.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return TransferTxResult{}, err
		}
	} else {
		// update to account balance.
		result.ToAccount, err = qtx.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return TransferTxResult{}, err
		}

		// update from account balance.
		result.FromAccount, err = qtx.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return TransferTxResult{}, err
		}
	}

	return result, tx.Commit(ctx)
}
