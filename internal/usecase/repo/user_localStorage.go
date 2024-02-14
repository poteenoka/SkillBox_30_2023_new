package repo

import (
	"context"
	"errors"
	"github.com/Skillbox_30_2023_new/internal/entity"
)

type UserLocalstorage struct {
	User map[string]*entity.User
}

func NewUserLocalstorage() *UserLocalstorage {
	return &UserLocalstorage{
		User: make(map[string]*entity.User),
	}
}

func (m *UserLocalstorage) CreateUser(ctx context.Context, user *entity.User) error {
	m.User[user.Name] = user
	return nil
}

func (m *UserLocalstorage) GetUser(ctx context.Context, name string) (*entity.User, error) {
	for key, _ := range m.User {
		if key == name {
			return m.User[key], nil
		}
	}
	return nil, errors.New("Пользователя не существует")
}

func (m *UserLocalstorage) UpdateUser(ctx context.Context, user *entity.User) error {
	for key, _ := range m.User {
		if key == user.Name {
			m.User[key] = user
			return nil
		}
	}
	return errors.New("Пользователя не существует")
}

func (m *UserLocalstorage) DeleteUser(ctx context.Context, id string) error {
	for key, value := range m.User {
		if value.ID == id {
			delete(m.User, key)
			return nil
		}
	}
	return errors.New("Пользователя не существует")
}

func (l *UserLocalstorage) MakeFriends(ctx context.Context, sourceID int, targetID int) error {
	return nil
}

func (l *UserLocalstorage) GetFriends(ctx context.Context, id int) (*entity.Userfriends, error) {
	return nil, nil
}
