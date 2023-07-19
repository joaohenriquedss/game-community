package entityGame

import (
	gameDtos "api/usecase/game/dtos"
)

type UserRepository interface {
	Insert(id string, name string, coverUrl string) error
	FindAll() ([]gameDtos.GamesDtoOutput, error)
}
