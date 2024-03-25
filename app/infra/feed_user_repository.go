package infra

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type FeedUserRepository interface {
	GetUserFeeds(userName string) ([]uuid.UUID, error)
	SetUserFeed(userName string, feedId uuid.UUID) error
}

type feedUserRepositoryImp struct {
	feedsPerUser map[string]map[uuid.UUID]interface{}

	lock *sync.RWMutex
}

func (f *feedUserRepositoryImp) GetUserFeeds(user string) ([]uuid.UUID, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	feedsIDs, exists := f.feedsPerUser[user]
	if !exists {
		return []uuid.UUID{}, fmt.Errorf("user %s does not have any feed", user)
	}
	var feedIDsValues []uuid.UUID
	for k, _ := range feedsIDs {
		feedIDsValues = append(feedIDsValues, k)
	}
	return feedIDsValues, nil
}

func (f *feedUserRepositoryImp) SetUserFeed(user string, feedId uuid.UUID) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	_, exists := f.feedsPerUser[user]
	if !exists {
		f.feedsPerUser[user] = make(map[uuid.UUID]interface{}, 0)
	}
	f.feedsPerUser[user][feedId] = nil
	return nil
}
