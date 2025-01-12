package db

import (
	"context"
	"testing"

	"github.com/akhushalani/rasa/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) Users {
	arg := CreateUserParams{
		Name:         pgtype.Text{String: "Test User", Valid: true},
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomString(32),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name.String, user.Name.String)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	fetchedUser, err := testQueries.GetUser(context.Background(), user.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedUser)

	require.Equal(t, user.UserID, fetchedUser.UserID)
	require.Equal(t, user.Name.String, fetchedUser.Name.String)
	require.Equal(t, user.Email, fetchedUser.Email)
	require.Equal(t, user.PasswordHash, fetchedUser.PasswordHash)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	arg := UpdateUserParams{
		UserID:       user.UserID,
		Name:         pgtype.Text{String: "Updated Name", Valid: true},
		Email:        "updated_" + user.Email,
		PasswordHash: "updated_" + user.PasswordHash,
	}

	err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)

	updatedUser, err := testQueries.GetUser(context.Background(), user.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, arg.Name.String, updatedUser.Name.String)
	require.Equal(t, arg.Email, updatedUser.Email)
	require.Equal(t, arg.PasswordHash, updatedUser.PasswordHash)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.UserID)
	require.NoError(t, err)

	// Verify user no longer exists
	fetchedUser, err := testQueries.GetUser(context.Background(), user.UserID)
	require.Error(t, err)
	require.Empty(t, fetchedUser)
}
