package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/maxgoover/rezonit-test-task/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)

	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Age, user2.Age)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, account := range users {
		require.NotEmpty(t, account)
	}
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:        user1.ID,
		FirstName: util.GenerateRandomString(6),
		LastName:  util.GenerateRandomString(10),
		Age:       util.GenerateRandomInt(1, 100),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.ID, user2.ID)
	require.Equal(t, arg.FirstName, user2.FirstName)
	require.Equal(t, arg.LastName, user2.LastName)
	require.Equal(t, arg.Age, user2.Age)
}

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		FirstName: util.GenerateRandomString(6),
		LastName:  util.GenerateRandomString(10),
		Age:       util.GenerateRandomInt(1, 100),
	}

	fmt.Printf("new user = %v", arg)
	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Age, user.Age)

	require.NotZero(t, user.ID)

	return user
}
