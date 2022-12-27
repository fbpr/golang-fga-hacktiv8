package main

import (
	"assignment-dua/database"
	"assignment-dua/router"
)

func main() {
	database.StartDB()
	r := router.StartServer()

	r.Run()
}
