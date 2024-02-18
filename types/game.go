package types

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type Game struct {
	players map[*websocket.Conn]*Player
}

func NewGame() *Game {
	return &Game{
		players: make(map[*websocket.Conn]*Player, 0),
	}
}

func (g *Game) AddPlayer(conn *websocket.Conn) {
	g.players[conn] = NewPlayer()
}

func (g *Game) RemovePlayer(conn *websocket.Conn) {
	delete(g.players, conn)
}

func (g *Game) UpdateState() {
	for _, player := range g.players {
		if player.IsMoving() {
			player.UpdatePosition()
		} else {
			break
		}
	}
}

func (g Game) GetCurrentState() string {
	gameDTO := GameToGameDTO(g)
	jsonState := gameDTO.ToJSON()
	fmt.Println(jsonState)
	return gameDTO.ToJSON()
}

func (g *Game) MovePlayerTowards(conn *websocket.Conn, position *Position) {
	g.players[conn].MoveTowards(position)
}