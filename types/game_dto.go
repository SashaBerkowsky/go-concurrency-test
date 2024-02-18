package types

import (
	"encoding/json"
)

type GameDTO struct {
	Players []PlayerDTO `json:"players"`
}

func GameToGameDTO(game Game) *GameDTO {
	playersDTO := make([]PlayerDTO, 0)
	i := 0

	for _, player := range(game.players) {
		playerDTO := player.ToDTO(i)
		playersDTO = append(playersDTO, *playerDTO)
		i += 1
	}

	return &GameDTO{
		Players: playersDTO,
	}
}

func (gameDTO GameDTO) ToJSON() string {
	gameJSON, err := json.Marshal(gameDTO)
    if err != nil {
		panic(err)
    }

	return string(gameJSON)
}
