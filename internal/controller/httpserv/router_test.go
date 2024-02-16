package httpserv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Skillbox_30_2023_new/internal/usecase/repo"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/Skillbox_30_2023_new/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repo.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	service := usecase.NewUserService(mockRepo)
	handler := NewHTTPHandler(service)

	handler.Rt.Post("/user", handler.CreateUser)
	handler.Rt.Get("/user/{name}", handler.GetUser)
	handler.Rt.Put("/user/{id}", handler.UpdateUser)
	handler.Rt.Delete("/user", handler.DeleteUser)
	handler.Rt.Post("/make_friends", handler.MakeFriends)
	handler.Rt.Get("/friends/{id}", handler.GetFriends)
	handler.Rt.Put("/user/age/{id}", handler.UpdateAge)

	fmt.Println("Test: CREATE USer")

	req, err := http.NewRequest("POST", "/user", bytes.NewBufferString(`{"name": "ivan", "age": 88}`))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	handler.CreateUser(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}
	var user entity.User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)

	assert.Equal(t, "ivan", user.Name)
	assert.Equal(t, 88, user.Age)

	fmt.Println("\n --------------", "Test Get USer", "-----------------")

	req, err = http.NewRequest("GET", "/user/ivan", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec = httptest.NewRecorder()

	handler.GetUser(rec, req)

	body, err = ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "ivan", user.Name)
	assert.Equal(t, 88, user.Age)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "ivan", user.Name)
	assert.Equal(t, 88, user.Age)

	fmt.Println("Test UPDATE USer")

	req, err = http.NewRequest("PUT", "/user/55", bytes.NewBufferString(`{"name": "ivan", "age": 31}`))
	if err != nil {
		t.Fatal(err)
	}

	rec = httptest.NewRecorder()

	//handler.UpdateUser(rec, req)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	assert.Equal(t, "ivan", user.Name)
	//  надо переделать на {name всесто ID.. или перед этим брать USER тфьу}assert.Equal(t, 31, user.Age)

}

/*
func TestGetUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repo.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().GetUser(gomock.Any(), "Ivan").Return(&entity.User{
		Name: "Ivan",
		Age:  88,
	}, nil).AnyTimes()

	service := usecase.NewUserService(mockRepo)
	handler := NewHTTPHandler(service)

	req, err := http.NewRequest("GET", "/user/Ivan", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	var user entity.User

	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, 30, user.Age)

	// Assert that the mock repository was called with the correct arguments.

	//ctrl.AssertCalled(t, "GetUser", gomock.Any(), "John Doe")
}

/*
func TestDeleteUser(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().DeleteUser(gomock.Any(), "1").Return(nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := NewHTTPHandler(service)

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
	mockRepo.mockRepo.AssertCalled(t, "DeleteUser", gomock.Any(), "1")
}

func TestMakeFriends(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().MakeFriends(gomock.Any(), 1, 2).Return(nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := NewHTTPHandler(service)

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
	mockRepo.mockRepo.AssertCalled(t, "MakeFriends", gomock.Any(), 1, 2)
}

func TestGetFriends(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().GetFriends(gomock.Any(), 1).Return(&entity.Userfriends{
		Friends: []int{2, 3},
	}, nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := NewHTTPHandler(service)

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
	mockRepo.mockRepo.AssertCalled(t, "GetFriends", gomock.Any(), 1)
}

func TestUpdateAge(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().GetUser(gomock.Any(), "John Doe").Return(&entity.User{
		Name: "John Doe",
		Age:  30,
	}, nil).AnyTimes()
	mockRepo.mockRepo.EXPECT().UpdateUser(gomock.Any(), &entity.User{
		Name: "John Doe",
		Age:  31,
	}).Return(nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := NewHTTPHandler(service)

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
	mockRepo.mockRepo.AssertCalled(t, "GetUser", gomock.Any(), "John Doe")
	mockRepo.mockRepo.AssertCalled(t, "UpdateUser", gomock.Any(), &entity.User{
		Name: "John Doe",
		Age:  31,
	})
}

func TestServRun(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

	// Create a new HTTP handler.
	handler := NewHTTPHandler(service)

	// Create a new HTTP server.
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Start the HTTP server.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			t.Fatal(err)
		}
	}()

	// Create a new HTTP client.
	client := &http.Client{}

	// Make a request to the HTTP server.
	resp, err := client.Get("http://localhost:8080/user")
	if err != nil {
		t.Fatal(err)
	}

	// Check the HTTP status code.
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(resp.Body)
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

	// Stop the HTTP server.
	server.Shutdown(context.Background())
}

func TestMain(m *testing.M) {
	// Create a new database connection.
	db, err := sql.Open("mssql", "server=localhost;user id=sa;password=your_password;port=1433;database=your_database")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create a new user repository.
	repo := repo.NewMSSQLUserRepository(db)

	// Create a new user service.
	service := usecase.NewUserService(repo)

	// Create a new HTTP handler.
	handler := NewHTTPHandler(service)

	// Create a new HTTP server.
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Start the HTTP server.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	// Run the tests.
	code := m.Run()

	// Stop the HTTP server.
	server.Shutdown(context.Background())

	// Exit the program.
	os.Exit(code)
}

/*
package httpserv

import (
	"bytes"
	"fmt"

	//"database/sql"
	"encoding/json"
	//"fmt"
	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/Skillbox_30_2023_new/internal/usecase"
	"github.com/Skillbox_30_2023_new/internal/usecase/repo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {

	mockRepo := new(repo.MockUserRepository)

	service := usecase.NewUserService(mockRepo)
	handler := NewHTTPHandler(service)
	req, err := http.NewRequest("POST", "/user", bytes.NewBufferString(`{"id": "1,"name": "Iam", "age": 50}`))
	if err != nil {
		t.Fatal(err)
	}
	// Create a new HTTP recorder.
	rec := httptest.NewRecorder()
	// Serve the HTTP request.
	handler.CreateUser(rec, req)

	// Check the HTTP status code.
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Read the HTTP response body.
	body, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodystring := string(body)
	fmt.Println("njdfsfsdf  :   ", bodystring)
	// Unmarshal the HTTP response body into a user.
	var user entity.User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}

	// Check the user's fields.
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, 30, user.Age)

	// Assert that the mock repository was called with the correct arguments.
	mockRepo.AssertCalled(t, "CreateUser", gomock.Any(), &user)
}

/*
func TestGetUser(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().GetUser(gomock.Any(), "John Doe").Return(&entity.User{
		Name: "John Doe",
		Age:  30,
	}, nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

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
	mockRepo.mockRepo.AssertCalled(t, "GetUser", gomock.Any(), "John Doe")
}

func TestUpdateUser(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().UpdateUser(gomock.Any(), &entity.User{
		Name: "Jane Doe",
		Age:  31,
	}).Return(nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

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
	mockRepo.mockRepo.AssertCalled(t, "UpdateUser", gomock.Any(), &entity.User{
		Name: "Jane Doe",
		Age:  31,
	})
}

func TestDeleteUser(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().DeleteUser(gomock.Any(), "1").Return(nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

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
	mockRepo.mockRepo.AssertCalled(t, "DeleteUser", gomock.Any(), "1")
}

func TestMakeFriends(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().MakeFriends(gomock.Any(), 1, 2).Return(nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

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
	mockRepo.mockRepo.AssertCalled(t, "MakeFriends", gomock.Any(), 1, 2)
}

func TestGetFriends(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().GetFriends(gomock.Any(), 1).Return(&entity.Userfriends{
		Friends: []int{2, 3},
	}, nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

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
	mockRepo.mockRepo.AssertCalled(t, "GetFriends", gomock.Any(), 1)
}

func TestUpdateAge(t *testing.T) {
	// Create a new mock user repository.
	mockRepo := NewMockUserRepository(t)

	// Set up the expected behavior of the mock repository.
	mockRepo.mockRepo.EXPECT().GetUser(gomock.Any(), "John Doe").Return(&entity.User{
		Name: "John Doe",
		Age:  30,
	}, nil).AnyTimes()
	mockRepo.mockRepo.EXPECT().UpdateUser(gomock.Any(), &entity.User{
		Name: "John Doe",
		Age:  31,
	}).Return(nil).AnyTimes()

	// Create a new user service.
	service := usecase.NewUserService(mockRepo)

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
	mockRepo.mockRepo.AssertCalled(t, "GetUser", gomock.Any(), "John Doe")
	mockRepo.mockRepo.AssertCalled(t, "UpdateUser", gomock.Any(), &entity.User{
		Name: "John Doe",
		Age:  31,
	})
}


func TestMain(m *testing.M) {
	// Create a new database connection.
	db, err := sql.Open("mssql", "server=localhost;user id=sa;password=your_password;port=1433;database=your_database")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create a new user repository.
	repo := repo.NewMSSQLUserRepository(db)

	// Create a new user service.
	service := usecase.NewUserService(repo)

	// Create a new HTTP handler.
	handler := NewHTTPHandler(service)

	// Create a new HTTP server.
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	// Start the HTTP server.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
	// Run the tests.
	code := m.Run()
	// Stop the HTTP server.
	server.Shutdown(context.Background())
	// Exit the program.
	os.Exit(code)
}
*/
