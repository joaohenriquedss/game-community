package user

import user "api/entity/entityUser"

type ProcessUser struct {
	UserRepository user.UserRepository
}

// & ponteiro para o objeto
// * mostrar o valor do objeto
func NewProcessUser(userRepository user.UserRepository) *ProcessUser {
	return &ProcessUser{
		UserRepository: userRepository,
	}
}
