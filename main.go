package main

import (
	// "net/http"

	"io"
	"os"
	"src/github.com/blackonyyx/cognizant/middlewares"
	"src/github.com/blackonyyx/cognizant/src/controller"
	"src/github.com/blackonyyx/cognizant/src/service"

	"github.com/gin-gonic/gin"
)

var (
	bookService service.BookService = service.New()
	bookController controller.BookController = controller.New(bookService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), )
	// server.Use()
	server.GET("/books", func (ctx *gin.Context) {
		ctx.JSON(200, bookController.FindAll())
	})
	server.POST("/save", func(ctx *gin.Context) {
		ctx.JSON(200, bookController.Save(ctx))
	})
	server.Run(":8080")
}
