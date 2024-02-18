package connection

import (
	"concurrency-test/types"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

const TICKER_TIME = time.Second

type Server struct {
	game 		 *types.Game
	clients		 map[*websocket.Conn]bool
	addClient    chan *websocket.Conn
	removeClient chan *websocket.Conn
	broadcast    chan []byte
	ticker		 *time.Ticker
}

func NewServer() *Server {
	return &Server {
		game: 		  types.NewGame(),
		clients: 	  make(map[*websocket.Conn]bool),
		addClient:    make(chan *websocket.Conn),
		removeClient: make(chan *websocket.Conn),
		broadcast:    make(chan []byte),
		ticker: time.NewTicker(TICKER_TIME),
	}
}

func (server *Server) handleWebSocket(conn *websocket.Conn) {
	server.addClient <- conn
	defer func() { server.removeClient <- conn }()
	for {
		var message string
		err := websocket.Message.Receive(conn, &message)
		if err != nil {
			log.Println("Error reading message from client:", err)
			break
		}
		// hacia donde va el personaje (click)
		playerInput := parsePlayerPosition(message)
		server.game.MovePlayerTowards(conn, playerInput)

		server.broadcast <- []byte(message)
	}
}

func parsePlayerPosition(positionJSON string) *types.Position {
	return types.NewPositionFromJSON(positionJSON)
}

func (server *Server) readLoop() {
	for {
		select {
		case client := <-server.addClient:
			server.game.AddPlayer(client)
			server.clients[client] = true
			log.Println("Client connected", client.RemoteAddr())
		case client := <-server.removeClient:
			server.game.RemovePlayer(client)
			delete(server.clients, client)
			log.Println("Client disconnected", client.RemoteAddr())
		case <- server.ticker.C:
			server.game.UpdateState()
			gameState := server.game.GetCurrentState()
			server.sendState(gameState)
		}
	}
}

func (server Server) sendState(gameState string) {
	for client := range server.clients {
		websocket.Message.Send(client, gameState)
	}
}

func (server Server) StartConnection() {
	http.Handle("/ws", websocket.Handler(server.handleWebSocket))

	// go server.broadcastState()
	go server.readLoop()

	log.Println("WebSocket server running on :3009")
	log.Fatal(http.ListenAndServe(":3009", nil))
}