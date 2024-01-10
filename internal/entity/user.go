package entity

import "github.com/Skillbox_30_2023_new/internal/usecase"

// User represents a user in the system.
type User struct {
	ID      string
	Name    string
	Age     int
	Friends []string
	Repo    usecase.UserRepository
}
