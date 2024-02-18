package types

import (
	"encoding/json"
)

type PositionDTO struct {
	X float32 `json:"x"`
	Z float32 `json:"z"`
}

func NewPositionFromJSON(jsonMessage string) *Position {
	var positionDTO PositionDTO = PositionDTO{}

	if err := json.Unmarshal([]byte(jsonMessage), &positionDTO); err != nil {
		panic(err)
	}

	return NewPosition(positionDTO.X, positionDTO.Z)
}
