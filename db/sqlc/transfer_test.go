package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	randomAmount := gofakeit.Int64()
	arg := CreateTransferParams{
		FromAccountID: pgtype.Int8{Int64: fromAccount.ID, Valid: true},
		ToAccountID:   pgtype.Int8{Int64: toAccount.ID, Valid: true},
		Amount:        randomAmount,
	}
	transfer, err := testStore.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	require.Equal(t, fromAccount.ID, transfer.FromAccountID.Int64)
	require.Equal(t, toAccount.ID, transfer.ToAccountID.Int64)
	require.Equal(t, randomAmount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testStore.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID.Int64, transfer2.FromAccountID.Int64)
	require.Equal(t, transfer1.ToAccountID.Int64, transfer2.ToAccountID.Int64)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	for range 10 {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  int32(5),
		Offset: int32(5),
	}
	transfers, err := testStore.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.NotEmpty(t, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.FromAccountID.Int64)
		require.NotZero(t, transfer.ToAccountID.Int64)
		require.NotZero(t, transfer.CreatedAt)
	}
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	newFromAccount := createRandomAccount(t)
	newToAccount := createRandomAccount(t)
	newAmount := gofakeit.Int64()
	arg := UpdateTransferParams{
		ID:            transfer1.ID,
		FromAccountID: pgtype.Int8{Int64: newFromAccount.ID, Valid: true},
		ToAccountID:   pgtype.Int8{Int64: newToAccount.ID, Valid: true},
		Amount:        newAmount,
	}
	transfer2, err := testStore.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.NotZero(t, transfer2.ID)
	require.Equal(t, transfer1.CreatedAt, transfer2.CreatedAt)
	require.Equal(t, newFromAccount.ID, transfer2.FromAccountID.Int64)
	require.Equal(t, newToAccount.ID, transfer2.ToAccountID.Int64)
	require.Equal(t, newAmount, transfer2.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	err := testStore.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testStore.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, transfer2)
}
