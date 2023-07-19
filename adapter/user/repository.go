package user

import (
	"api/usecase/user/dtos"
	"database/sql"
	"time"
)

type UserDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type users struct {
	db *sql.DB
}

func NewUserRepositoryDB(db *sql.DB) *users {
	return &users{db: db}
}

func (db *users) Insert(id string, name string, email string, password string) error {
	_, err := db.db.Exec("INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", id, name, email, password, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (db *users) FindAll() ([]dtos.UserDtoOutput, error) {
	rows, err := db.db.Query("SELECT id,name,email FROM users")
	if err != nil {
		return nil, err
	}

	var users []dtos.UserDtoOutput
	for rows.Next() {
		var id string
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)

		if err != nil {
			return nil, err
		}
		user := dtos.UserDtoOutput{
			Id:    id,
			Name:  name,
			Email: email,
		}
		users = append(users, user)
	}

	return users, nil
}
