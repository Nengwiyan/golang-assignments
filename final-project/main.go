package main

import (
	"final-project/database"
	"final-project/routers"
)

var PORT = ":8080"

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(PORT)
}
