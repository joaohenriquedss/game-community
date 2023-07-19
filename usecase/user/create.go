package user

import (
	user "api/entity/entityUser"
	"api/usecase/user/dtos"
)

func (processUser *ProcessUser) Execute(userCreateDto dtos.UserCreateDto) (dtos.UserDtoOutput, error) {
	input := user.NewUser()
	input.Name = userCreateDto.Name
	input.Email = userCreateDto.Email
	input.Password = userCreateDto.Password
	input.PasswordConfirm = userCreateDto.PasswordConfirm
	err := input.IsValid()
	if err != nil {
		return processUser.errorUser(err)
	}
	return processUser.approveUser(input)
}

func (processUser *ProcessUser) approveUser(input *user.User) (dtos.UserDtoOutput, error) {
	err := processUser.UserRepository.Insert(input.Id, input.Name, input.Email, input.Password)
	if err != nil {
		return processUser.errorUser(err)
	}
	output := dtos.UserDtoOutput{
		Id:    input.Id,
		Name:  input.Name,
		Email: input.Email,
	}
	return output, nil
}

func (p *ProcessUser) errorUser(err error) (dtos.UserDtoOutput, error) {
	return dtos.UserDtoOutput{}, err
}
