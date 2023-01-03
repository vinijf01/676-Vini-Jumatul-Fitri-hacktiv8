package main

import (
	"securing-with-jwt/database"
	"securing-with-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
