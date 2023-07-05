package game

import "fmt"

type Game struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Console       string `json:"console"`
	Score         int    `json:"score"`
	Review        string `json:"review"`
	Status        string `json:"status"`
	Genre         string `json:"genre"`
	CoverUrl      string `json:"coverUrl"`
	MinutesPlayed int    `json:"minutesPlayed"`
}

func IsValid() {
	fmt.Println("criando a interface  do")
}
