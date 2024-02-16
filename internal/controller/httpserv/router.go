package httpserv

import (
	"fmt"
	"github.com/Skillbox_30_2023_new/cmd/config"
	"github.com/Skillbox_30_2023_new/internal/usecase"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"regexp"
)

type HTTPHandler struct {
	Service *usecase.UserService
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("путь", r.URL.Path)
	var routeUSer = regexp.MustCompile(`/user.*`)
	var routeFriend = regexp.MustCompile(`/friends.*`)

	switch r.Method {
	case "POST":
		switch r.URL.Path {
		case "/user":
			fmt.Println("Post User....")
			h.CreateUser(w, r)
		case "/make_friends":
			h.MakeFriends(w, r)
		default:
			http.NotFound(w, r)
		}
	case "GET":
		switch cmd := r.URL.Path; {
		case routeUSer.MatchString(cmd):
			fmt.Println("Берем пользователя")
			h.GetUser(w, r)
		case routeFriend.MatchString(cmd):
			h.GetFriends(w, r)
		default:
			http.NotFound(w, r)

		}
	case "PUT":
		switch r.URL.Path {
		case "/user/{id}":
			h.UpdateUser(w, r)
		case "/user/age/{id}":
			h.UpdateAge(w, r)
		default:
			http.NotFound(w, r)
		}
	case "DELETE":
		switch r.URL.Path {
		case "/user":
			h.DeleteUser(w, r)
		default:
			http.NotFound(w, r)
		}
	default:
		http.NotFound(w, r)
	}
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
