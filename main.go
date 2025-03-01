package main

import (
	// "net/http"

	"fmt"
	"io"
	"net/http"
	"os"
	"src/github.com/blackonyyx/cognizant/middlewares"
	"src/github.com/blackonyyx/cognizant/src/controller"
	"src/github.com/blackonyyx/cognizant/src/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
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
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())
	// server.Use()
	server.GET("/books", func (ctx *gin.Context) {
		books := bookController.FindAll()
		fmt.Println(books)
		ctx.JSON(200, books)
	})
	server.POST("/save", func(ctx *gin.Context) {
		res, err := bookController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, res.String())
		}
	})
	server.Run(":8080")
}
