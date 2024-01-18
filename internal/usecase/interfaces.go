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
}

// бизнес логика, авторизации и прочие действия...
type UseCase interface {
	SignInuser(ctx context.Context, user *entity.User) error
	signOut(ctx context.Context, user *entity.User) error
}
