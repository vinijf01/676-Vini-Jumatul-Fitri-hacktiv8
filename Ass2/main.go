package main

import (
	"ass2/databases"
	"ass2/routers"
)

func main() {
	databases.StartDB()
	// fmt.Print("test db")
	routers.GetRouter().Run(":3000")
}
