package user

import (
	dtos "api/usecase/user/dtos"
)

func (processUser *ProcessUser) FindAll() ([]dtos.UserDtoOutput, error) {
	users, err := processUser.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
