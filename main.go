package main

import (
	"gin-h8/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
