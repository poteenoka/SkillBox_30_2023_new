package usecase

import (
	"context"
	"github.com/Skillbox_30_2023_new/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id string) error
	MakeFriends(ctx context.Context, sourceID int, targetID int) error
	GetFriends(ctx context.Context, id int) (*entity.Userfriends, error)
}
