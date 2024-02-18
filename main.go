package main

import (
	"concurrency-test/connection"
)

func main() {
	server := connection.NewServer()

	server.StartConnection()
}