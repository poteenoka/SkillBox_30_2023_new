package repo

import (
	"context"
	"errors"
	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserLocalstorage_GetUser(t *testing.T) {
	storage := NewUserLocalstorage()

	user := &entity.User{
		Name: "Иван Васильевич",
		Age:  30,
	}
	err := storage.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	retrievedUser, err := storage.GetUser(context.Background(), user.Name)

	assert.NoError(t, err)
	assert.Equal(t, user, retrievedUser)
}

func TestUserLocalstorage_UpdateUser(t *testing.T) {
	storage := NewUserLocalstorage()
	user := &entity.User{
		Name: "Иван Васильевич",
		Age:  30,
	}
	err := storage.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	user.Age = 31
	err = storage.UpdateUser(context.Background(), user)
	assert.NoError(t, err)
	retrievedUser, err := storage.GetUser(context.Background(), user.Name)

	assert.NoError(t, err)
	assert.Equal(t, user.Age, retrievedUser.Age)
}

func TestUserLocalstorage_DeleteUser(t *testing.T) {
	storage := NewUserLocalstorage()
	user := &entity.User{
		Name: "John Doe",
		Age:  30,
	}

	err := storage.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	retrievedUser, err := storage.GetUser(context.Background(), user.Name)

	assert.Nil(t, retrievedUser)
	assert.ErrorIs(t, err, errors.New("Пользователя не существует"))

	err = storage.DeleteUser(context.Background(), user.ID)
	assert.NoError(t, err)

}
