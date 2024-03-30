package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var userrepo UserRepository

func TestMain(t *testing.M) {
	userrepo = NewUserRepository()
}
func TestCreateUser(t *testing.T) {
	err := userrepo.CreateUser("ben")
	require.NoError(t, err)
	err = userrepo.CreateUser("gal")
	require.NoError(t, err)
	users, err := userrepo.GetAllUsers()
	require.NoError(t, err)
	assert.Equal(t, 2, len(users))
}

func TestDeleteUser(t *testing.T) {
	userrepo := NewUserRepository()
	err := userrepo.CreateUser("ben")
	require.NoError(t, err)
	err = userrepo.CreateUser("gal")
	require.NoError(t, err)
	err = userrepo.DeleteUser("ben")
	require.NoError(t, err)
	assert.Equal(t, 1, len(userrepo.users))
}

func TestUpdateUser(t *testing.T) {
	userrepo := NewUserRepository()
	err := userrepo.CreateUser("ben")
	require.NoError(t, err)
	err = userrepo.CreateUser("gal")
	require.NoError(t, err)
	err = userrepo.UpdateUser("gal", "new_gal")
	require.NoError(t, err)
	_, err = userrepo.GetUser("new_gal")
	require.NoError(t, err)
}
