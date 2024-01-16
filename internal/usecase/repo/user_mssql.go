package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/google/uuid"
)

type MSSQLUserRepository struct {
	db *sql.DB
}

func NewMSSQLUserRepository(db *sql.DB) *MSSQLUserRepository {
	return &MSSQLUserRepository{
		db: db,
	}
}

// Create creates a new user in the repository.
func (r *MSSQLUserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	user.ID = uuid.New().String()
	fmt.Println(user)
	_, err := r.db.Exec("INSERT INTO users (name, age) VALUES ( @name, @age)",
		sql.Named("name", user.Name),
		sql.Named("age", user.Age),
	)
	return err
}

// Get gets a user from the repository by their ID.
func (r *MSSQLUserRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id, name, age FROM users WHERE id = @id", sql.Named("id", id))
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user in the repository.
func (r MSSQLUserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	_, err := r.db.Exec("UPDATE users SET name = @name, age = @age WHERE id = @id",
		sql.Named("id", user.ID),
		sql.Named("name", user.Name),
		sql.Named("age", user.Age),
	)
	return err
}

// Delete deletes a user from the repository.
func (r MSSQLUserRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = @id", sql.Named("id", id))
	return err
}
