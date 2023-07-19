package entityUser

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type User struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func NewUser() *User {
	id := uuid.New()
	return &User{
		Id: id.String(),
	}
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

func isValidName(name string) bool {
	return len(name) >= 3
}

func isValidPassword(password string) bool {
	return len(password) >= 8
}

func isValidPasswordConfirm(password string, passwordConfirm string) bool {
	return password == passwordConfirm
}

func (user *User) IsValid() error {
	valid_email := isValidEmail(user.Email)
	if !valid_email {
		return errors.New("invalid email")
	}
	valid_name := isValidName(user.Name)
	if !valid_name {
		return errors.New("invalid name")
	}
	valid_password := isValidPassword(user.Password)
	if !valid_password {
		return errors.New("invalid password")
	}
	valid_password_confirm := isValidPasswordConfirm(user.Password, user.PasswordConfirm)
	if !valid_password_confirm {
		return errors.New("invalid password confirm")
	}
	return nil
}
