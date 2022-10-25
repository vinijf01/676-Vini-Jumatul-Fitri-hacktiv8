package main

import (
	"mygram/databases"
	"mygram/routers"
)

func main() {
	databases.StartDB()
	routers.GetRouter().Run(":3000")
}
