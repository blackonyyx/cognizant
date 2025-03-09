package main

import (
	// "log"
	"src/github.com/blackonyyx/cognizant/src"

	// "github.com/gin-gonic/gin"
)

// "net/http"



func main() {
	server := src.SetupRouter()
	// log.Println("Hello there")
	server.Run(":3000")
}
