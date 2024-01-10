package usecase

import (
	"context"
	"github.com/Skillbox_30_2023_new/internal/entity"
)

// UserRepository is an interface for managing users.
type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}
