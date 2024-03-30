package handlers

import (
	"fmt"
	"sort"
	"sync"

	"github.com/bohana3/twitter-like-go/app/infra"
	"github.com/bohana3/twitter-like-go/app/models"
)

type Follower interface {
	// FollowUser adds a follower to a following user (following)
	FollowUser(followerName string, followingName string) error
	// UnfollowUser removes a follower from a following user (following)
	UnfollowUser(followerName string, followingName string) error
	// GetMutualFollowers get a list of users that follow both userName1 and userName2
	GetMutualFollowers(userName1 string, userName2 string) ([]*models.User, error)
	// GetTopInfluencers returns the top 'n' users with the highest number of followers
	GetTopInfluencers(n int) ([]models.User, error)
}

type FollowerHandler struct {
	users     infra.UserRepository
	followers infra.FollowerRepository
	lock      *sync.RWMutex
}

func NewFollowerHandler(users infra.UserRepository, followers infra.FollowerRepository) *FollowerHandler {
	return &FollowerHandler{
		users:     users,
		followers: followers,
		lock:      &sync.RWMutex{},
	}
}

func (f *FollowerHandler) FollowUser(followerName string, followingName string) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	_, err := f.users.GetUser(followerName)
	if err != nil {
		return err
	}
	_, err = f.users.GetUser(followingName)
	if err != nil {
		return err
	}

	err = f.followers.SetFollower(followerName, followingName)
	if err != nil {
		return err
	}
	return nil
}

func (f *FollowerHandler) UnfollowUser(followerName string, followingName string) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	_, err := f.users.GetUser(followerName)
	if err != nil {
		return err
	}
	_, err = f.users.GetUser(followingName)
	if err != nil {
		return err
	}

	followed, err := f.followers.GetFollowers(followingName)
	if err != nil {
		return err
	}

	if len(followed.Followers) == 0 {
		return fmt.Errorf("user %s does not have any follower", followerName)
	}

	_, exist := followed.Followers[followerName]
	if !exist {
		return fmt.Errorf("follower %s does not follows %s", followerName, followingName)
	}

	err = f.followers.RemoveFollower(followerName, followingName)
	if err != nil {
		return err
	}
	return nil
}

// GetMutualFollowers returns mutual followers – users that follow both userName1 and userName2
func (f *FollowerHandler) GetMutualFollowers(userName1 string, userName2 string) ([]*models.User, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	following1, err := f.followers.GetFollowers(userName1)
	if err != nil {
		return []*models.User{}, err
	}
	if len(following1.Followers) == 0 {
		return []*models.User{}, nil
	}

	following2, err := f.followers.GetFollowers(userName2)
	if err != nil {
		return []*models.User{}, err
	}
	if len(following2.Followers) == 0 {
		return []*models.User{}, nil
	}

	var mutualFollower []string
	for follower1 := range following1.Followers {
		_, exist := following2.Followers[follower1]
		if exist {
			mutualFollower = append(mutualFollower, follower1)
		}
	}

	var users []*models.User
	for _, follower := range mutualFollower {
		user, err := f.users.GetUser(follower)
		if err != nil {
			return []*models.User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (f *FollowerHandler) GetTopInfluencers(n int) ([]models.User, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	followings, err := f.followers.GetAllFollowers()
	if err != nil {
		return []models.User{}, err
	}

	sort.Slice(followings, func(i, j int) bool {
		return len(followings[i].Followers) > len(followings[j].Followers)
	})

	var users []models.User
	for i := 0; i < n; i++ {
		user, err := f.users.GetUser(followings[i].Name)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, *user)
	}
	return users, nil
}
