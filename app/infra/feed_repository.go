package infra

import (
	"fmt"
	"sync"

	"github.com/bohana3/twitter-like-go/app/models"
	"github.com/google/uuid"
)

type FeedRepository interface {
	GetFeed(id uuid.UUID) (*models.Tweet, error)
	GetFeeds(id []uuid.UUID) ([]*models.Tweet, error)
	SetFeed(tweetContent string) (*models.Tweet, error)
}

type feedRepositoryImp struct {
	feeds map[uuid.UUID]*models.Tweet

	lock *sync.RWMutex
}

func NewFeedRepository() *feedRepositoryImp {
	return &feedRepositoryImp{
		feeds: make(map[uuid.UUID]*models.Tweet, 0),
		lock:  &sync.RWMutex{},
	}
}

func (f *feedRepositoryImp) GetFeed(feedID uuid.UUID) (*models.Tweet, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	tweet, exist := f.feeds[feedID]
	if !exist {
		return &models.Tweet{}, fmt.Errorf("feed %s does not exist", feedID)
	}
	return tweet, nil
}

func (f *feedRepositoryImp) GetFeeds(feedIDs []uuid.UUID) ([]*models.Tweet, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	var tweets []*models.Tweet
	for _, feedID := range feedIDs {
		tweet, exist := f.feeds[feedID]
		if !exist {
			return []*models.Tweet{}, fmt.Errorf("feed %s does not exist", feedID)
		}
		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

func (f *feedRepositoryImp) SetFeed(tweetContent string) (*models.Tweet, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	tweet, err := models.NewTweet(tweetContent)
	if err != nil {
		return &models.Tweet{}, err
	}
	f.feeds[tweet.ID] = &tweet
	return &tweet, nil
}
