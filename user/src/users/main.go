package main

import (
	"os"
)


func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	server := UserServer()
	server.Run(":" + port)
}