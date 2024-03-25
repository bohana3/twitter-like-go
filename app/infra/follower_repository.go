package infra

import (
	"fmt"
	"sync"

	"github.com/bohana3/twitter-like-go/app/models"
)

type FollowerRepository interface {
	GetAllFollowers() ([]*models.Followed, error)
	// Get followers that follow a specific user
	GetFollowers(followingName string) (*models.Followed, error)
	// Get users that are followed by a specific user
	GetFolloweds(followerName string) ([]*models.Followed, error)
	SetFollower(followerName string, followingName string) error
	RemoveFollower(followerName string, followingName string) error
}

type followerRepositoryImpl struct {
	followed map[string]models.Followed // key: Following name  // value: Following

	lock *sync.RWMutex
}

func (f *followerRepositoryImpl) GetFollowers(followingName string) (*models.Followed, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	followings, exist := f.followed[followingName]
	if !exist {
		return &models.Followed{}, nil
	}
	return &followings, nil
}

func (f *followerRepositoryImpl) GetFolloweds(followerName string) ([]*models.Followed, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	var followings []*models.Followed
	for _, f := range f.followed {
		_, exist := f.Followers[followerName]
		if exist {
			followings = append(followings, &f)
		}
	}
	return followings, nil
}

func (f *followerRepositoryImpl) SetFollower(followingName string, followerName string) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	_, exist := f.followed[followingName]
	if !exist {
		f.followed[followingName] = models.NewFollowing(followingName)
	}
	f.followed[followingName].Followers[followerName] = nil
	return nil
}

func (f *followerRepositoryImpl) RemoveFollower(followingName string, followerName string) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	_, exist := f.followed[followingName]
	if !exist {
		return fmt.Errorf("user %s does not have any follower", followingName)
	}
	delete(f.followed[followingName].Followers, followerName)
	return nil
}

func (f *followerRepositoryImpl) GetAllFollowers() ([]*models.Followed, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	var followings []*models.Followed
	for _, f := range f.followed {
		followings = append(followings, &f)
	}

	return followings, nil
}
