package infra

import (
	"fmt"
	"sync"

	"github.com/bohana3/twitter-like-go/app/models"
)

type FollowerRepository interface {
	GetAllFollowers() ([]*models.FollowedUser, error)
	// Get followers that follow a specific user
	GetFollowers(followingName string) (*models.FollowedUser, error)
	// Get users that are followed by a specific user
	GetFolloweds(followerName string) ([]*models.FollowedUser, error)
	SetFollower(followerName string, followingName string) error
	RemoveFollower(followerName string, followingName string) error
}

type followerRepositoryImpl struct {
	followedUsers map[string]*models.FollowedUser // key: Following name  // value: Following

	lock *sync.RWMutex
}

func NewFollowerRepository() *followerRepositoryImpl {
	return &followerRepositoryImpl{
		followedUsers: make(map[string]*models.FollowedUser, 0),
		lock:          &sync.RWMutex{},
	}
}

func (f *followerRepositoryImpl) GetFollowers(followingName string) (*models.FollowedUser, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	followings, exist := f.followedUsers[followingName]
	if !exist {
		return &models.FollowedUser{}, nil
	}
	return followings, nil
}

func (f *followerRepositoryImpl) GetFolloweds(followerName string) ([]*models.FollowedUser, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	var followings []*models.FollowedUser
	for _, f := range f.followedUsers {
		_, exist := f.Followers[followerName]
		if exist {
			followings = append(followings, f)
		}
	}
	return followings, nil
}

func (f *followerRepositoryImpl) SetFollower(followerName string, followingName string) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	_, exist := f.followedUsers[followingName]
	if !exist {
		f.followedUsers[followingName] = models.NewFollowing(followingName)
	}
	f.followedUsers[followingName].Followers[followerName] = nil
	return nil
}

func (f *followerRepositoryImpl) RemoveFollower(followerName string, followingName string) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	_, exist := f.followedUsers[followingName]
	if !exist {
		return fmt.Errorf("user %s does not have any follower", followingName)
	}
	delete(f.followedUsers[followingName].Followers, followerName)
	if len(f.followedUsers[followingName].Followers) == 0 {
		delete(f.followedUsers, followingName)
	}
	return nil
}

func (f *followerRepositoryImpl) GetAllFollowers() ([]*models.FollowedUser, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	var followings []*models.FollowedUser
	for _, f := range f.followedUsers {
		followings = append(followings, f)
	}

	return followings, nil
}
