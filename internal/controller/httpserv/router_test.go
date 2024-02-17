package httpserv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Skillbox_30_2023_new/internal/usecase/repo"
	"github.com/go-chi/chi"
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

	fmt.Println("Test: CREATE USer")

	req, err := http.NewRequest("POST", "/user", bytes.NewBufferString(`{"id": "1", "name": "ivan", "age": 88}`))
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
	//-----------------

	fmt.Println("\n --------------", "Test Get USer", "-----------------")

	mockRepo.EXPECT().GetUser(gomock.Any(), "ivan").Return(&entity.User{
		Name: "ivan",
		Age:  88,
	}, nil).AnyTimes()

	w := httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/user/{name}", nil)
	if err != nil {
		t.Fatal(err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "ivan")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	chi.URLParam(req, "ivan")

	handler.GetUser(w, req)

	body, err = ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ivan", user.Name)
	assert.Equal(t, 88, user.Age)

	//-------------------------

	fmt.Println("Test UPDATE USer")
	mockRepo.EXPECT().UpdateUser(gomock.Any(), &entity.User{
		ID:   "1",
		Name: "ivan",
		Age:  18,
	}).Return(nil).AnyTimes()

	req, err = http.NewRequest("PUT", "/user/{id}", bytes.NewBufferString(`{"name": "ivan", "age": 18}`))
	if err != nil {
		t.Fatal(err)
	}

	h := httptest.NewRecorder()
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	chi.URLParam(req, "id")

	handler.UpdateUser(h, req)

	assert.Equal(t, http.StatusOK, h.Code)
	assert.Equal(t, "ivan", user.Name)

	//--------------------------
	fmt.Println("delete user: ...")
	mockRepo.EXPECT().DeleteUser(gomock.Any(), "1").Return(nil).AnyTimes()
	req, err = http.NewRequest("DELETE", "/user/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	//rctx = chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	chi.URLParam(req, "1")
	handler.DeleteUser(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	body, err = ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "{\"message\":\"user deleted\"}\n", string(body))

}
