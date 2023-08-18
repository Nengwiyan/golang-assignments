package main

import (
	"assignment-project/database"
	"assignment-project/routers"
)

var (
	PORT = ":8080"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(PORT)
}
