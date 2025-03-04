package main

import "src/github.com/blackonyyx/cognizant/src"

// "net/http"


func main() {
	server := src.SetupRouter()
	server.Run(":3000")
}
