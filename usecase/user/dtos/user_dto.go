package dtos

type UserCreateDto struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type UserDtoOutput struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
