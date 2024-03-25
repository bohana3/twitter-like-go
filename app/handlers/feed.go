package handlers

import (
	"sort"

	"github.com/bohana3/twitter-like-go/app/infra"
	"github.com/bohana3/twitter-like-go/app/models"
)

type Feed interface {
	GetUserFeed(userName string) ([]models.Tweet, error)
	PostTweet(userName string, tweetContent string) error
}

type FeedHandler struct {
	users     infra.UserRepository
	followers infra.FollowerRepository
	feeds     infra.FeedRepository
	userfeeds infra.FeedUserRepository
}

func (f *FeedHandler) PostTweet(userName string, tweetContent string) error {
	user, err := f.users.GetUser(userName)
	if err != nil {
		return err
	}
	tweet, err := f.feeds.SetFeed(tweetContent)
	if err != nil {
		return err
	}
	err = f.userfeeds.SetUserFeed(user.Name, tweet.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetUserFeed returns the list of tweets from other users that this user follows
// remark: sorted by most recent
func (f *FeedHandler) GetUserFeed(userName string) ([]models.Tweet, error) {
	followeds, err := f.followers.GetFolloweds(userName)
	if err != nil {
		return []models.Tweet{}, err
	}
	var tweets []models.Tweet
	for _, followed := range followeds {
		userfeedIDs, err := f.userfeeds.GetUserFeeds(followed.Name)
		if err != nil {
			return []models.Tweet{}, err
		}
		feeds, err := f.feeds.GetFeeds(userfeedIDs)
		if err != nil {
			return []models.Tweet{}, err
		}
		for _, feed := range feeds {
			tweets = append(tweets, *feed)
		}
	}

	// sort by most recent order
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].CreationTime > tweets[j].CreationTime
	})

	return tweets, nil
}
