package token

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	key := gofakeit.Password(true, false, false, false, false, 32)
	maker, err := NewPasetoMaker(key)
	require.NoError(t, err)

	username := gofakeit.Password(true, false, false, false, false, 16)
	duration := time.Minute

	issutedAt := time.Now()
	expiredAt := issutedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issutedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, issutedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	key := gofakeit.Password(true, false, false, false, false, 32)
	maker, err := NewPasetoMaker(key)
	require.NoError(t, err)

	username := gofakeit.Password(true, false, false, false, false, 16)

	token, err := maker.CreateToken(username, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
