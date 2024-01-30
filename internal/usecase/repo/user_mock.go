package repo

import (
	"context"
	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUser(ctx context.Context, name string) (*entity.User, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) MakeFriends(ctx context.Context, sourceID int, targetID int) error {
	args := m.Called(ctx, sourceID, targetID)
	return args.Error(0)
}

func (m *MockUserRepository) GetFriends(ctx context.Context, id int) (*entity.Userfriends, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Userfriends), args.Error(1)
}
