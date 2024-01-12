package main

import (
	"database/sql"
	"github.com/Skillbox_30_2023_new/internal/controller/httpserv"
	"github.com/Skillbox_30_2023_new/internal/usecase"
	"github.com/Skillbox_30_2023_new/internal/usecase/repo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlserver", "server=localhost;user id=sa;password=YourStrong@Password;port=1433;database=user_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repo.NewMSSQLUserRepository(db)
	service := usecase.NewUserService(repo)
	handler := httpserv.NewHTTPHandler(service)

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
