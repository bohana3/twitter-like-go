package models

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID
	Name string
}

func NewUser(name string) User {
	newUuid, _ := uuid.NewUUID()
	return User{
		ID:   newUuid,
		Name: name,
	}
}

// copy ctor
func NewUserFromUser(name string, user User) User {
	newUuid, _ := uuid.NewUUID()
	return User{
		ID:   newUuid,
		Name: name,
	}
}
