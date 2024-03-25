package models

import "github.com/google/uuid"

const MaxTweetLength = 300

type Tweet struct {
	ID           uuid.UUID
	Content      *string
	CreationTime uint64
}

func NewTweet(content string) (Tweet, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return Tweet{}, err
	}
	if len(content) > MaxTweetLength { // limit Tweet size
		content = content[:MaxTweetLength]
	}
	return Tweet{
		ID:      uuid,
		Content: &content,
	}, nil
}

type FeedRepository interface {
	GetUserFeed(userName string) ([]Tweet, error)
	GetMutualFollowers(userName1 string, userName2 string) ([]*User, error)
}
