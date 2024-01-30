package httpserv

import (
	"github.com/Skillbox_30_2023_new/cmd/config"
	"github.com/Skillbox_30_2023_new/internal/usecase"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

type HTTPHandler struct {
	Service *usecase.UserService
}

func (h *HTTPHandler) ServeHTTP() {

}

func NewHTTPHandler(service *usecase.UserService) *HTTPHandler {
	return &HTTPHandler{
		Service: service,
	}
}

func ServRun(service *usecase.UserService, cfg *config.Config) {
	handler := NewHTTPHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/user", handler.CreateUser)
	r.Get("/user/{name}", handler.GetUser)
	r.Put("/user/{id}", handler.UpdateUser)
	r.Delete("/user", handler.DeleteUser)
	r.Post("/make_friends", handler.MakeFriends)
	r.Get("/friends/{id}", handler.GetFriends)
	r.Put("/user/age/{id}", handler.UpdateAge)

	port := cfg.HTTP.Port
	port = ":" + port
	log.Fatal(http.ListenAndServe(port, r))
}
