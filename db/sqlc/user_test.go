package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/phantom-sage/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	password := gofakeit.Password(true, false, false, false, false, 16)
	hashedPassword, err := util.HashPasswrod(password)
	require.NoError(t, err)

	randomUsername := gofakeit.Password(true, false, false, false, false, 16)
	randomHashedPassword := hashedPassword
	randomFullname := gofakeit.HackerNoun()
	randomEmail := gofakeit.Email()

	arg := CreateUserParams{
		Username:       randomUsername,
		HashedPassword: randomHashedPassword,
		FullName:       randomFullname,
		Email:          randomEmail,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEmpty(t, user.Username)
	require.NotEmpty(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	require.Equal(t, user.HashedPassword, randomHashedPassword)
	require.Equal(t, user.FullName, randomFullname)
	require.Equal(t, user.Email, randomEmail)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
