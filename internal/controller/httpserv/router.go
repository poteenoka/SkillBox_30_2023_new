package httpserv

import (
	"encoding/json"
	"fmt"
	"github.com/Skillbox_30_2023_new/cmd/config"
	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/Skillbox_30_2023_new/internal/usecase"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type HTTPHandler struct {
	Service *usecase.UserService
	Rt      *chi.Mux
}

func (*HTTPHandler) ServeHTTP(http.ResponseWriter, *http.Request) {

}

func NewHTTPHandler(service *usecase.UserService) *HTTPHandler {
	return &HTTPHandler{
		Service: service,
		Rt:      chi.NewRouter(),
	}
}

func ServRun(service *usecase.UserService, cfg *config.Config) {
	handler := NewHTTPHandler(service)
	//	r := chi.NewRouter()

	handler.Rt.Post("/user", handler.CreateUser)
	handler.Rt.Get("/user/{name}", handler.GetUser)
	handler.Rt.Put("/user/{id}", handler.UpdateUser)
	handler.Rt.Delete("/user", handler.DeleteUser)
	handler.Rt.Post("/make_friends", handler.MakeFriends)
	handler.Rt.Get("/friends/{id}", handler.GetFriends)
	handler.Rt.Put("/user/age/{id}", handler.UpdateAge)

	port := cfg.HTTP.Port
	port = ":" + port
	log.Fatal(http.ListenAndServe(port, handler.Rt))
}

func (h *HTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("CreateUSer:   ", user.Age, user.ID, user.Name)
	if err := h.Service.CreateUser(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *HTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	name := chi.URLParam(r, "name")
	fmt.Println("Имя: ", name)
	user, err := h.Service.GetUser(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

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

func (h *HTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.Service.DeleteUser(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user deleted"})
}

func (h *HTTPHandler) MakeFriends(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SourceID int `json:"source_id"`
		TargetID int `json:"target_id"`
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

func (h *HTTPHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idstr)
	friends, err := h.Service.GetFriends(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(friends)
}

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
