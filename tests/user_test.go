package tests

import (
	"context"
	db "github.com/ngenohkevin/ark-realtors/db/sqlc"
	"github.com/ngenohkevin/ark-realtors/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) db.User {

	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := db.CreateUserParams{
		ID:             utils.GenerateRandomUserID(),
		Username:       utils.RandomUsername(),
		FullName:       utils.RandomFullName(),
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

}

func TestUpdateUserOnlyFullName(t *testing.T) {
	oldUser := createRandomUser(t)

	newFullName := utils.RandomFullName()
	updatedUser, err := testStore.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       oldUser.ID,
		FullName: utils.NullStrings(newFullName),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, newFullName, updatedUser.FullName)
	require.Equal(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Role, updatedUser.Role)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)

	newEmail := utils.RandomEmail()
	updatedUser, err := testStore.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:    oldUser.ID,
		Email: utils.NullStrings(newEmail),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Role, updatedUser.Role)
}

func TestUpdateUserOnlyRole(t *testing.T) {
	oldUser := createRandomUser(t)

	newRole := utils.RandomRole()
	updatedUser, err := testStore.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:   oldUser.ID,
		Role: utils.NullStrings(newRole),
	})

	require.NoError(t, err)
	//require.NotEqual(t, oldUser.Role, updatedUser.Role)
	require.Equal(t, newRole, updatedUser.Role)
	require.Equal(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	newPassword := utils.RandomString(6)
	hashedPassword, err := utils.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testStore.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:             oldUser.ID,
		HashedPassword: utils.NullStrings(hashedPassword),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, hashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.Role, updatedUser.Role)
}

func TestUpdateUsernameOnly(t *testing.T) {
	oldUser := createRandomUser(t)

	newUsername := utils.RandomUsername()
	updatedUser, err := testStore.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       oldUser.ID,
		Username: utils.NullStrings(newUsername),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, newUsername, updatedUser.Username)
	require.Equal(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Role, updatedUser.Role)
}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := createRandomUser(t)

	newUsername := utils.RandomUsername()
	newFullName := utils.RandomFullName()
	newEmail := utils.RandomEmail()
	newRole := utils.RandomRole()
	newPassword := utils.RandomString(6)
	hashedPassword, err := utils.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testStore.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:             oldUser.ID,
		Username:       utils.NullStrings(newUsername),
		FullName:       utils.NullStrings(newFullName),
		Email:          utils.NullStrings(newEmail),
		HashedPassword: utils.NullStrings(hashedPassword),
		Role:           utils.NullStrings(newRole),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, newUsername, updatedUser.Username)
	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, newFullName, updatedUser.FullName)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	//require.NotEqual(t, oldUser.Role, updatedUser.Role)
	require.Equal(t, newRole, updatedUser.Role)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, hashedPassword, updatedUser.HashedPassword)
}
