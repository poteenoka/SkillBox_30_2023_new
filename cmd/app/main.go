package main

import (
	"database/sql"
	"fmt"
	"github.com/Skillbox_30_2023_new/cmd/config"
	"github.com/Skillbox_30_2023_new/internal/controller/httpserv"
	"github.com/Skillbox_30_2023_new/internal/usecase"
	"github.com/Skillbox_30_2023_new/internal/usecase/repo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	connMssql := fmt.Sprintf("server=localhost;user id=%s;password=%s;port=1433;database=user_db", cfg.MSSQL.User, cfg.MSSQL.Password)
	db, err := sql.Open("sqlserver", connMssql)
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

	port := cfg.HTTP.Port
	log.Fatal(http.ListenAndServe(port, r))
}
