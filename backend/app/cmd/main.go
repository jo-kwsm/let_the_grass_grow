package main

import (
	"grass_backend/db"
	"grass_backend/server"
)

func main() {
	db.Init()
	server.Init()
}
