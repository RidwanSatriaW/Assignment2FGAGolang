package main

import (
	"assignment2/database"
	"assignment2/routers"
)

func main() {
	database.StartDB()
	r:=routers.StartServer()
	r.Run(":8080")
}
