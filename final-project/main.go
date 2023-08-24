package main

import (
	"final-project/database"
	"final-project/routers"
)

var (
	PORT = ":7070"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(PORT)
}
