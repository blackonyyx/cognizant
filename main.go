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
	"src/github.com/blackonyyx/cognizant/src/service/book"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	bookService book.BookService = book.New()
	bookController controller.BookController = controller.NewBookController(bookService)
	loanService service.LoanService = service.New(bookService)
	loanController controller.LoanController = controller.NewLoanController(loanService)

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
	server.POST("/add" , func (ctx *gin.Context) {

		res, err := bookController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, res.String())
		}
	})
	server.GET("/books", func (ctx *gin.Context) {
		books := bookController.FindAll()
		fmt.Println(books)
		ctx.JSON(200, books)
	})

	server.GET("/read", func (ctx *gin.Context) { // /:key
		read, err := bookController.GetContent(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, read)
		}

	})

	server.GET("/search", func(ctx *gin.Context) {
		books, err := bookController.GetBooks(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, books)
		}
	})

	server.GET("/loan", func(ctx *gin.Context) {
		resp, err := loanController.GetLoanReceipt(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, resp)
		}
	})

	server.POST("/return", func(ctx *gin.Context) {
		resp, err := loanController.ReturnBooks(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, resp)
		}
	})
	server.POST("/extend", func(ctx *gin.Context) {
		resp, err := loanController.ExtendLoan(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, resp)
		}
	})
	server.Run(":3000")
}
