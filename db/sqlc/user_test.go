package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	//  username, hashed_password, full_name, email
	randomUsername := gofakeit.Password(true, false, false, false, false, 16)
	randomHashedPassword := gofakeit.Password(true, false, false, false, false, 16)
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
	require.NotEmpty(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)

	require.Equal(t, user.HashedPassword, randomHashedPassword)
	require.Equal(t, user.FullName, randomFullname)
	require.Equal(t, user.Email, randomEmail)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
