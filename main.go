package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type GameDtoInput struct {
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

type GameDtoOutput struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	CoverUrl string `json:"coverUrl"`
}

func main() {
	var games = []GameDtoOutput{
		{Id: 1, Name: "The Last of Us", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 2, Name: "Super Mario World", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 3, Name: "Red Dead Redemption 2", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co2vus.jpg"},
		{Id: 4, Name: "The Witcher 3: Wild Hunt", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1vbj.jpg"},
		{Id: 5, Name: "Grand Theft Auto V", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1vu9.jpg"},
		{Id: 6, Name: "The Legend of Zelda: Breath of the Wild", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1vhg.jpg"},
		{Id: 7, Name: "Minecraft", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1vj8.jpg"},
		{Id: 8, Name: "Portal 2", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1w2s.jpg"},
		{Id: 9, Name: "Dark Souls III", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co20s7.jpg"},
		{Id: 10, Name: "Super Mario Odyssey", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1w2i.jpg"},
		{Id: 11, Name: "The Elder Scrolls V: Skyrim", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 12, Name: "The Last of Us Part II", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 13, Name: "God of War", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 14, Name: "Uncharted 4: A Thief's End", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 15, Name: "The Legend of Zelda: Ocarina of Time", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 16, Name: "Hollow Knight", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 17, Name: "Bloodborne", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 18, Name: "The Legend of Zelda: A Link to the Past", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 19, Name: "Super Metroid", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
		{Id: 20, Name: "Mass Effect 2", CoverUrl: "https://images.igdb.com/igdb/image/upload/t_cover_big/co1v5i.jpg"},
	}

	http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {

		jsonResponse, err := json.Marshal(games)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonResponse)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
