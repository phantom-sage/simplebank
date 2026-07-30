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

func createRandomEntry(t *testing.T) Entry {
	randomAccount := createRandomAccount(t)
	randomAmount := gofakeit.Int64()
	arg := CreateEntryParams{
		AccountID: pgtype.Int8{Int64: randomAccount.ID, Valid: true},
		Amount:    randomAmount,
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt.Time)

	require.Equal(t, randomAccount.ID, entry.AccountID.Int64)
	require.Equal(t, randomAmount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID.Int64, entry2.AccountID.Int64)
	require.Equal(t, entry1.Amount, entry2.Amount)

	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)
}

func TestListEntries(t *testing.T) {
	for range 10 {
		createRandomEntry(t)
	}

	arg := ListEntrysParams{
		Limit:  int32(5),
		Offset: int32(5),
	}
	entries, err := testQueries.ListEntrys(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.NotEmpty(t, entry.Amount)
		require.NotZero(t, entry.ID)
		require.NotZero(t, entry.AccountID.Int64)
		require.NotZero(t, entry.CreatedAt.Time)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	newAccount := createRandomAccount(t)
	newAmount := gofakeit.Int64()
	arg := UpdateEntryParams{
		ID:        entry1.ID,
		AccountID: pgtype.Int8{Int64: newAccount.ID, Valid: true},
		Amount:    newAmount,
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.NotZero(t, entry2.ID)
	require.Equal(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time)
	require.Equal(t, newAccount.ID, entry2.AccountID.Int64)
	require.Equal(t, newAmount, entry2.Amount)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, entry2)
}
