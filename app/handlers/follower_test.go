package handlers

import (
	"testing"

	"github.com/bohana3/twitter-like-go/app/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var followerrepo infra.FollowerRepository
var userrepo infra.UserRepository
var followerHandler *FollowerHandler

func TestBaseFollower(t *testing.T) {
	userrepo = infra.NewUserRepository()
	_ = userrepo.CreateUser("ben")

	followerrepo = infra.NewFollowerRepository()
	followerHandler = NewFollowerHandler(userrepo, followerrepo)
}
func TestFollowUser(t *testing.T) {
	TestBaseFollower(t)

	err := userrepo.CreateUser("elon musk")
	require.NoError(t, err)
	err = followerHandler.FollowUser("ben", "elon musk")
	require.NoError(t, err)
	followers, err := followerrepo.GetAllFollowers()
	require.NoError(t, err)
	assert.Equal(t, 1, len(followers))
}

func TestUnFollowUser(t *testing.T) {
	TestBaseFollower(t)

	err := userrepo.CreateUser("elon musk")
	require.NoError(t, err)
	err = followerHandler.FollowUser("ben", "elon musk")
	require.NoError(t, err)
	err = followerHandler.UnfollowUser("ben", "elon musk")
	require.NoError(t, err)
	followers, err := followerrepo.GetAllFollowers()
	require.NoError(t, err)
	assert.Equal(t, 0, len(followers))
}

func TestGetMutualFollowers(t *testing.T) {
	TestBaseFollower(t)

	err := userrepo.CreateUser("elon musk")
	require.NoError(t, err)
	err = userrepo.CreateUser("trump")
	require.NoError(t, err)
	err = followerHandler.FollowUser("ben", "elon musk")
	require.NoError(t, err)
	err = followerHandler.FollowUser("ben", "trump")
	require.NoError(t, err)
	mutualUsers, err := followerHandler.GetMutualFollowers("elon musk", "trump")
	require.NoError(t, err)
	assert.Equal(t, 1, len(mutualUsers))
	assert.Equal(t, "ben", mutualUsers[0].Name)
}

func TestGetTopInfluencers(t *testing.T) {
	TestBaseFollower(t)

	_ = userrepo.CreateUser("gal")
	_ = userrepo.CreateUser("shaun")
	_ = userrepo.CreateUser("elon musk")
	_ = userrepo.CreateUser("trump")
	_ = userrepo.CreateUser("byte-byte-go")

	err := followerHandler.FollowUser("ben", "elon musk")
	require.NoError(t, err)
	err = followerHandler.FollowUser("ben", "trump")
	require.NoError(t, err)
	err = followerHandler.FollowUser("ben", "byte-byte-go")
	require.NoError(t, err)

	err = followerHandler.FollowUser("gal", "elon musk")
	require.NoError(t, err)
	err = followerHandler.FollowUser("gal", "byte-byte-go")
	require.NoError(t, err)

	err = followerHandler.FollowUser("shaun", "byte-byte-go")
	require.NoError(t, err)

	users, err := followerHandler.GetTopInfluencers(3)
	require.NoError(t, err)
	assert.Equal(t, 3, len(users))
	assert.Equal(t, "byte-byte-go", users[0].Name)
	assert.Equal(t, "elon musk", users[1].Name)
	assert.Equal(t, "trump", users[2].Name)
}
