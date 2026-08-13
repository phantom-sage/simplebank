package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomSession(t *testing.T) Session {
	user := createRandomUser(t)

	id, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	refreshToken := gofakeit.Password(true, false, true, true, false, 16)
	userAgent := gofakeit.HackerNoun()
	clientIP := "127.0.0.1"
	expiresAt := time.Now()

	arg := CreateSessionParams{
		ID:           id,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
		ClientIp:     clientIP,
		IsBlocked:    true,
		ExpiresAt:    expiresAt,
	}
	session, err := testStore.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.NotZero(t, session.CreatedAt)
	require.NotZero(t, session.ExpiresAt)

	require.Equal(t, session.ID, id)
	require.Equal(t, session.Username, user.Username)
	require.Equal(t, session.RefreshToken, refreshToken)
	require.Equal(t, session.UserAgent, userAgent)
	require.Equal(t, session.ClientIp, clientIP)

	require.True(t, session.IsBlocked)

	require.WithinDuration(t, session.ExpiresAt, expiresAt, time.Second)

	return session
}

func TestCreateSession(t *testing.T) {
	createRandomSession(t)
}

func TestGetSession(t *testing.T) {
	session1 := createRandomSession(t)
	session2, err := testStore.GetSession(context.Background(), session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, session1.ID, session2.ID)
	require.Equal(t, session1.Username, session2.Username)
	require.Equal(t, session1.RefreshToken, session2.RefreshToken)
	require.Equal(t, session1.UserAgent, session2.UserAgent)
	require.Equal(t, session1.ClientIp, session2.ClientIp)
	require.Equal(t, session1.IsBlocked, session2.IsBlocked)

	require.WithinDuration(t, session1.CreatedAt, session2.CreatedAt, time.Second)
	require.WithinDuration(t, session1.ExpiresAt, session2.ExpiresAt, time.Second)
}
