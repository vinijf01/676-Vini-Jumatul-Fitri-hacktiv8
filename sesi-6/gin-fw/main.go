package main

import "gin-fw/routers"

func main() {
	routers.StartServer().Run(":8080")
}
