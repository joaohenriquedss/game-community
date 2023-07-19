package dtos

type GameCreateDtos struct {
	Name     string `json:"name"`
	CoverUrl string `json:"coverUrl"`
}

type GamesDtoOutput struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	CoverUrl string `json:"coverUrl"`
}
