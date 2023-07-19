package entityUser

import (
	userDtos "api/usecase/user/dtos"
)

type UserRepository interface {
	Insert(id string, name string, email string, password string) error
	FindAll() ([]userDtos.UserDtoOutput, error)
}
