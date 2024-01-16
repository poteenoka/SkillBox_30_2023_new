package usecase

import (
	"context"
	"github.com/Skillbox_30_2023_new/internal/entity"
)

// UserService is a service for managing users.
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CreateUser creates a
// new user.
func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	return s.repo.CreateUser(ctx, user)
}

// GetUser gets a user by their ID.
func (s *UserService) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return s.repo.GetUser(ctx, id)
}

// UpdateUser updates a user.
func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	return s.repo.UpdateUser(ctx, user)
}

// DeleteUser deletes a user.
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}

// MakeFriends makes two users friends.
func (s *UserService) MakeFriends(ctx context.Context, sourceID, targetID string) error {
	sourceUser, err := s.repo.GetUser(ctx, sourceID)
	if err != nil {
		return err
	}
	targetUser, err := s.repo.GetUser(ctx, targetID)
	if err != nil {
		return err
	}
	sourceUser.Friends = append(sourceUser.Friends, targetID)
	targetUser.Friends = append(targetUser.Friends, sourceID)
	if err := s.repo.UpdateUser(ctx, sourceUser); err != nil {
		return err
	}
	if err := s.repo.UpdateUser(ctx, targetUser); err != nil {
		return err
	}
	return nil
}

// GetFriends gets all of a user's friends.
func (s *UserService) GetFriends(ctx context.Context, id string) ([]*entity.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	friends := make([]*entity.User, len(user.Friends))
	for i, friendID := range user.Friends {
		friend, err := s.repo.GetUser(ctx, friendID)
		if err != nil {
			return nil, err
		}
		friends[i] = friend
	}
	return friends, nil
}

// UpdateAge updates a user's age.
func (s *UserService) UpdateAge(ctx context.Context, id string, newAge int) error {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return err
	}
	user.Age = newAge
	return s.repo.UpdateUser(ctx, user)
}
