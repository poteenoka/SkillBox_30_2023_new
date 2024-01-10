package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/Skillbox_30_2023_new/internal/usecase"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// MSSQLUserRepository is an implementation of UserRepository that uses MSSQL as the storage backend.
type MSSQLUserRepository struct {
	db *sql.DB
}

// NewMSSQLUserRepository creates a new MSSQLUserRepository.
func NewMSSQLUserRepository(db *sql.DB) *MSSQLUserRepository {
	return &MSSQLUserRepository{
		db: db,
	}
}

// Create creates a new user in the repository.
func (r *MSSQLUserRepository) Create(ctx context.Context, user *entity.User) error {
	user.ID = uuid.New().String()
	_, err := r.db.Exec("INSERT INTO users (id, name, age) VALUES (@id, @name, @age)",
		sql.Named("id", user.ID),
		sql.Named("name", user.Name),
		sql.Named("age", user.Age),
	)
	return err
}

// Get gets a user from the repository by their ID.
func (r *MSSQLUserRepository) Get(ctx context.Context, id string) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id, name, age FROM users WHERE id = @id", sql.Named("id", id))
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user in the repository.
func (r *MSSQLUserRepository) Update(ctx context.Context, user *entity.User) error {
	_, err := r.db.Exec("UPDATE users SET name = @name, age = @age WHERE id = @id",
		sql.Named("id", user.ID),
		sql.Named("name", user.Name),
		sql.Named("age", user.Age),
	)
	return err
}

// Delete deletes a user from the repository.
func (r *MSSQLUserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = @id", sql.Named("id", id))
	return err
}

// UserService is a service for managing users.
type UserService struct {
	Repo usecase.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo usecase.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

// CreateUser creates a
// new user.
func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	return s.Repo.Create(ctx, user)
}

// GetUser gets a user by their ID.
func (s *UserService) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return s.Repo.Get(ctx, id)
}

// UpdateUser updates a user.
func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	return s.Repo.Update(ctx, user)
}

// DeleteUser deletes a user.
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}

// MakeFriends makes two users friends.
func (s *UserService) MakeFriends(ctx context.Context, sourceID, targetID string) error {
	sourceUser, err := s.Repo.Get(ctx, sourceID)
	if err != nil {
		return err
	}
	targetUser, err := s.Repo.Get(ctx, targetID)
	if err != nil {
		return err
	}
	sourceUser.Friends = append(sourceUser.Friends, targetID)
	targetUser.Friends = append(targetUser.Friends, sourceID)
	if err := s.Repo.Update(ctx, sourceUser); err != nil {
		return err
	}
	if err := s.Repo.Update(ctx, targetUser); err != nil {
		return err
	}
	return nil
}

// GetFriends gets all of a user's friends.
func (s *UserService) GetFriends(ctx context.Context, id string) ([]*entity.User, error) {
	user, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	friends := make([]*entity.User, len(user.Friends))
	for i, friendID := range user.Friends {
		friend, err := s.Repo.Get(ctx, friendID)
		if err != nil {
			return nil, err
		}
		friends[i] = friend
	}
	return friends, nil
}

// UpdateAge updates a user's age.
func (s *UserService) UpdateAge(ctx context.Context, id string, newAge int) error {
	user, err := s.Repo.Get(ctx, id)
	if err != nil {
		return err
	}
	user.Age = newAge
	return s.Repo.Update(ctx, user)
}

// HTTPHandler is a HTTP handler for the UserService.
type HTTPHandler struct {
	Service *UserService
}

// NewHTTPHandler creates a new HTTPHandler.
func NewHTTPHandler(service *UserService) *HTTPHandler {
	return &HTTPHandler{
		Service: service,
	}
}

// CreateUser creates a new user.
func (h *HTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateUser(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser gets a user by their ID.
func (h *HTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.Service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// UpdateUser updates a user.
func (h *HTTPHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = id
	if err := h.Service.UpdateUser(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// DeleteUser deletes a user.
func (h *HTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.Service.DeleteUser(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user deleted"})
}

// MakeFriends makes two users friends.
func (h *HTTPHandler) MakeFriends(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SourceID string `json:"source_id"`
		TargetID string `json:"target_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.MakeFriends(r.Context(), req.SourceID, req.TargetID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "users are now friends"})
}

// GetFriends gets all of a user's friends.
func (h *HTTPHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	friends, err := h.Service.GetFriends(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(friends)
}

// UpdateAge updates a user's age.
func (h *HTTPHandler) UpdateAge(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req struct {
		NewAge int `json:"new_age"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.UpdateAge(r.Context(), id, req.NewAge); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user's age updated"})
}

func main() {
	db, err := sql.Open("sqlserver", "server=localhost;user id=sa;password=YourStrong@Password;port=1433;database=user_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewMSSQLUserRepository(db)
	service := NewUserService(repo)
	handler := NewHTTPHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/user", handler.CreateUser)
	r.Get("/user/{id}", handler.GetUser)
	r.Put("/user/{id}", handler.UpdateUser)
	r.Delete("/user", handler.DeleteUser)
	r.Post("/make_friends", handler.MakeFriends)
	r.Get("/friends/{id}", handler.GetFriends)
	r.Put("/user/{id}/age", handler.UpdateAge)

	log.Fatal(http.ListenAndServe(":8080", r))
}
