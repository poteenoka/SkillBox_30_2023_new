package httpserv

import (
	"encoding/json"
	"fmt"
	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

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
