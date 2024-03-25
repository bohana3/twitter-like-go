package infra

import (
	"fmt"
	"sync"

	"github.com/bohana3/twitter-like-go/app/models"
)

type UserRepository interface {
	CreateUser(name string) error
	GetUser(name string) (*models.User, error)
	UpdateUser(oldName string, newName string) error
	DeleteUser(name string) error
}

type userRepositoryImpl struct {
	users map[string]*models.User

	lock *sync.RWMutex
}

func (u *userRepositoryImpl) CreateUser(name string) error {
	u.lock.Lock()
	defer u.lock.Unlock()

	_, exist := u.users[name]
	if exist {
		return fmt.Errorf("user %s already exists", name)
	}
	newUser := models.NewUser(name)
	u.users[name] = &newUser
	return nil
}

func (u *userRepositoryImpl) GetUser(name string) (*models.User, error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	user, exist := u.users[name]
	if !exist {
		return &models.User{}, fmt.Errorf("user %s does not exists", name)
	}
	return user, nil
}

func (u *userRepositoryImpl) UpdateUser(oldName string, newName string) error {
	u.lock.Lock()
	defer u.lock.Unlock()

	user, exist := u.users[oldName]
	if !exist {
		return fmt.Errorf("user %s does not exists", oldName)
	}
	newUser := models.NewUserFromUser(newName, *user)
	u.users[newName] = &newUser
	delete(u.users, oldName)
	return nil
}

func (u *userRepositoryImpl) DeleteUser(name string) error {
	u.lock.Lock()
	defer u.lock.Unlock()

	_, exist := u.users[name]
	if !exist {
		return fmt.Errorf("user %s does not exists", name)
	}
	delete(u.users, name)
	return nil
}
