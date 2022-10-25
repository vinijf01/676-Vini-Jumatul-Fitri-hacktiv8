package main

import (
	"ass3/controllers"
	"ass3/helpers"
	"ass3/services"
	"net/http"
	"time"
)

func main() {
	go autoUpdate()
	http.HandleFunc("/", controllers.DataController)

	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}
}

func autoUpdate() {
	data := helpers.GetData()
	for range time.Tick(15 * time.Second) {
		services.UpdateJSON(data)
	}
}
