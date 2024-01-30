package main

import (
	"database/sql"
	"fmt"
	"github.com/Skillbox_30_2023_new/cmd/config"
	"github.com/Skillbox_30_2023_new/internal/controller/httpserv"
	"github.com/Skillbox_30_2023_new/internal/usecase"
	"github.com/Skillbox_30_2023_new/internal/usecase/repo"
	_ "github.com/microsoft/go-mssqldb"
	"log"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	connMssql := fmt.Sprintf("server=localhost;user id=%s;password=%s;port=1433;database=%s", cfg.MSSQL.User, cfg.MSSQL.Password, cfg.MSSQL.DatabaseName)
	db, err := sql.Open("mssql", connMssql)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	repoNew := repo.NewMSSQLUserRepository(db)
	service := usecase.NewUserService(repoNew)
	httpserv.ServRun(service, cfg)
}
