package main

import (
	"pertemuan-sembilan/database"
	routers "pertemuan-sembilan/routers"
)

var (
	PORT = ":8080"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(PORT)
}
