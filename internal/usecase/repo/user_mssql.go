package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Skillbox_30_2023_new/internal/entity"
	"github.com/google/uuid"
	_ "github.com/microsoft/go-mssqldb"
	"log"
)

type MSSQLUserRepository struct {
	db *sql.DB
}

func NewMSSQLUserRepository(db *sql.DB) *MSSQLUserRepository {
	return &MSSQLUserRepository{
		db: db,
	}
}

func (r *MSSQLUserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	user.ID = uuid.New().String()
	//	fmt.Println("полезли в базу:....", user.Name, user.Age)
	err := r.db.Ping()
	query := fmt.Sprintf("select top 1 id from users where name = '%s'", user.Name)
	fmt.Println(query)
	row, _ := r.db.Query(query)
	var userid int
	for row.Next() {
		if err = row.Scan(&userid); err == nil {
			err = fmt.Errorf("пользователь %s уже заведен в системе", user.Name)
			return err
		}
	}

	_, err = r.db.Exec(fmt.Sprintf("INSERT INTO users (name, age) VALUES ( '%s', %d)", user.Name, user.Age))
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (r *MSSQLUserRepository) GetUser(ctx context.Context, name string) (*entity.User, error) {
	//idint, _ := strconv.Atoi(id)
	row := r.db.QueryRow("SELECT id, name, age FROM users WHERE name = ?", name)
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r MSSQLUserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	_, err := r.db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", user.Name, user.Age, user.ID)
	return err
}

func (r MSSQLUserRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = @id", sql.Named("id", id))
	return err
}
