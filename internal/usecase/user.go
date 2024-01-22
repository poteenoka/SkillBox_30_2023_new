package usecase

import (
	"context"
	"github.com/Skillbox_30_2023_new/internal/entity"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	return s.repo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) MakeFriends(ctx context.Context, sourceID, targetID int) error {
	return s.repo.MakeFriends(ctx, sourceID, targetID)
}

func (s *UserService) GetFriends(ctx context.Context, id int) (*entity.Userfriends, error) {

	return s.repo.GetFriends(ctx, id)
}

func (s *UserService) UpdateAge(ctx context.Context, id string, newAge int) error {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return err
	}
	user.Age = newAge
	return s.repo.UpdateUser(ctx, user)
}
