package usecase

import (
	"bytes"
	"encoding/json"
	"github.com/Skillbox_30_2023_new/internal/controller/httpserv"
	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/Skillbox_30_2023_new/internal/usecase/repo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {

	mockRepo := new(repo.MockUserRepository)
	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)
	service := NewUserService(mockRepo)
	handler := httpserv.NewHTTPHandler(service)

	// Create a new HTTP request.
	req, err := http.NewRequest("POST", "/user", bytes.NewBufferString(`{"name": "John Doe", "age": 30}`))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder.
	rec := httptest.NewRecorder()

	// Serve the HTTP request.
	handler.ServeHTTP(rec, req)

	// Check the HTTP status code.
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal the HTTP response body into a user.
	var user entity.User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}

	// Check the user's fields.
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, 30, user.Age)

	// Assert that the mock repository was called with the correct arguments.
	mockRepo.AssertCalled(t, "CreateUser", mock.Anything, &user)
}

func TestGetUser(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := new(repo.MockUserRepository)

	// Set up the expected behavior of the mock repository.
	mockRepo.On("GetUser", mock.Anything, "John Doe").Return(&entity.User{
		Name: "John Doe",
		Age:  30,
	}, nil)

	// Create a new user service.
	service := NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := httpserv.NewHTTPHandler(service)

	// Create a new HTTP request.
	req, err := http.NewRequest("GET", "/user/John Doe", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder.
	rec := httptest.NewRecorder()

	// Serve the HTTP request.
	handler.ServeHTTP(rec, req)

	// Check the HTTP status code.
	assert.Equal(t, http.StatusOK, rec.Code)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal the HTTP response body into a user.
	var user entity.User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}

	// Check the user's fields.
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, 30, user.Age)

	// Assert that the mock repository was called with the correct arguments.
	mockRepo.AssertCalled(t, "GetUser", mock.Anything, "John Doe")
}

func TestUpdateUser(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := new(repo.MockUserRepository)

	// Set up the expected behavior of the mock repository.
	mockRepo.On("UpdateUser", mock.Anything, &entity.User{
		Name: "Jane Doe",
		Age:  31,
	}).Return(nil)

	// Create a new user service.
	service := NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := httpserv.NewHTTPHandler(service)

	// Create a new HTTP request.
	req, err := http.NewRequest("PUT", "/user/1", bytes.NewBufferString(`{"name": "Jane Doe", "age": 31}`))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder.
	rec := httptest.NewRecorder()

	// Serve the HTTP request.
	handler.ServeHTTP(rec, req)

	// Check the HTTP status code.
	assert.Equal(t, http.StatusOK, rec.Code)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal the HTTP response body into a user.
	var user entity.User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}

	// Check the user's fields.
	assert.Equal(t, "Jane Doe", user.Name)
	assert.Equal(t, 31, user.Age)

	// Assert that the mock repository was called with the correct arguments.
	mockRepo.AssertCalled(t, "UpdateUser", mock.Anything, &entity.User{
		Name: "Jane Doe",
		Age:  31,
	})
}

func TestDeleteUser(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := new(repo.MockUserRepository)

	// Set up the expected behavior of the mock repository.
	mockRepo.On("DeleteUser", mock.Anything, "1").Return(nil)

	// Create a new user service.
	service := NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := httpserv.NewHTTPHandler(service)

	// Create a new HTTP request.
	req, err := http.NewRequest("DELETE", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder.
	rec := httptest.NewRecorder()

	// Serve the HTTP request.
	handler.ServeHTTP(rec, req)

	// Check the HTTP status code.
	assert.Equal(t, http.StatusOK, rec.Code)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the HTTP response body.
	assert.Equal(t, `{"message":"user deleted"}`, string(body))

	// Assert that the mock repository was called with the correct arguments.
	mockRepo.AssertCalled(t, "DeleteUser", mock.Anything, "1")
}

func TestMakeFriends(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := new(repo.MockUserRepository)

	// Set up the expected behavior of the mock repository.
	mockRepo.On("MakeFriends", mock.Anything, 1, 2).Return(nil)

	// Create a new user service.
	service := NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := httpserv.NewHTTPHandler(service)

	// Create a new HTTP request.
	req, err := http.NewRequest("POST", "/make_friends", bytes.NewBufferString(`{"source_id": 1, "target_id": 2}`))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder.
	rec := httptest.NewRecorder()

	// Serve the HTTP request.
	handler.ServeHTTP(rec, req)

	// Check the HTTP status code.
	assert.Equal(t, http.StatusOK, rec.Code)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the HTTP response body.
	assert.Equal(t, `{"message":"users are now friends"}`, string(body))

	// Assert that the mock repository was called with the correct arguments.
	mockRepo.AssertCalled(t, "MakeFriends", mock.Anything, 1, 2)
}

func TestGetFriends(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := new(repo.MockUserRepository)

	// Set up the expected behavior of the mock repository.
	mockRepo.On("GetFriends", mock.Anything, 1).Return(&entity.Userfriends{
		Friends: []int{2, 3},
	}, nil)

	// Create a new user service.
	service := NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := httpserv.NewHTTPHandler(service)

	// Create a new HTTP request.
	req, err := http.NewRequest("GET", "/friends/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder.
	rec := httptest.NewRecorder()

	// Serve the HTTP request.
	handler.ServeHTTP(rec, req)

	// Check the HTTP status code.
	assert.Equal(t, http.StatusOK, rec.Code)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal the HTTP response body into a slice of users.
	var friends entity.Userfriends
	if err := json.Unmarshal(body, &friends); err != nil {
		t.Fatal(err)
	}

	// Check the friends' fields.
	assert.Equal(t, 2, len(friends.Friends))
	assert.Equal(t, 2, friends.Friends[0])
	assert.Equal(t, 3, friends.Friends[1])

	// Assert that the mock repository was called with the correct arguments.
	mockRepo.AssertCalled(t, "GetFriends", mock.Anything, 1)
}

func TestUpdateAge(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := new(repo.MockUserRepository)

	// Set up the expected behavior of the mock repository.
	mockRepo.On("GetUser", mock.Anything, "John Doe").Return(&entity.User{
		Name: "John Doe",
		Age:  30,
	}, nil)
	mockRepo.On("UpdateUser", mock.Anything, &entity.User{
		Name: "John Doe",
		Age:  31,
	}).Return(nil)

	// Create a new user service.
	service := NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := httpserv.NewHTTPHandler(service)

	// Create a new HTTP request.
	req, err := http.NewRequest("PUT", "/user/age/1", bytes.NewBufferString(`{"new_age": 31}`))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder.
	rec := httptest.NewRecorder()

	// Serve the HTTP request.
	handler.ServeHTTP(rec, req)

	// Check the HTTP status code.
	assert.Equal(t, http.StatusOK, rec.Code)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the HTTP response body.
	assert.Equal(t, `{"message":"user's age updated"}`, string(body))

	// Assert that the mock repository was called with the correct arguments.
	mockRepo.AssertCalled(t, "GetUser", mock.Anything, "John Doe")
	mockRepo.AssertCalled(t, "UpdateUser", mock.Anything, &entity.User{
		Name: "John Doe",
		Age:  31,
	})
}
