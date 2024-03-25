package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	TweetIds []uuid.UUID
}

func NewUser(name string) User {
	newUuid, _ := uuid.NewUUID()
	return User{
		ID:       newUuid,
		Name:     name,
		TweetIds: make([]uuid.UUID, 0),
	}
}

// copy ctor
func NewUserFromUser(name string, user User) User {
	newUuid, _ := uuid.NewUUID()
	return User{
		ID:       newUuid,
		Name:     name,
		TweetIds: user.TweetIds,
	}
}
