package handlers

import (
	"testing"

	"github.com/bohana3/twitter-like-go/app/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var feedHandler *FeedHandler

func TestBaseFeed(t *testing.T) {
	userrepo := infra.NewUserRepository()
	_ = userrepo.CreateUser("ben")
	_ = userrepo.CreateUser("trump")
	_ = userrepo.CreateUser("elon musk")

	feedrepo = infra.NewFeedRepository()
	feeduserrepo := infra.NewFeedUserRepository()
	followerrepo = infra.NewFollowerRepository()

	feedHandler = NewFeedHandler(userrepo, followerrepo, feedrepo, feeduserrepo)
	followerHandler = NewFollowerHandler(userrepo, followerrepo)
}

func TestPostTweet(t *testing.T) {
	TestBaseFeed(t)

	tweet, err := feedHandler.PostTweet("ben", "this is ben first tweet")
	require.NoError(t, err)

	feed, err := feedrepo.GetFeed(tweet.ID)
	require.NoError(t, err)
	require.Equal(t, "this is ben first tweet", *feed.Content)
}

func TestGetUserFeed(t *testing.T) {
	TestBaseFeed(t)

	_, err := feedHandler.PostTweet("elon musk", "I have a breaking idea to improve Tweeter")
	require.NoError(t, err)
	_, err = feedHandler.PostTweet("trump", "I will be the next president of USA")
	require.NoError(t, err)

	err = followerHandler.FollowUser("ben", "trump")
	require.NoError(t, err)
	err = followerHandler.FollowUser("ben", "elon musk")
	require.NoError(t, err)
	feeds, err := feedHandler.GetUserFeed("ben")
	require.NoError(t, err)
	assert.Equal(t, 2, len(feeds))
}
