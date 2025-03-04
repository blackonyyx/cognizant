package src

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"src/github.com/blackonyyx/cognizant/src/controller"
	"src/github.com/blackonyyx/cognizant/src/errormsg"
	"src/github.com/blackonyyx/cognizant/src/middlewares"
	"src/github.com/blackonyyx/cognizant/src/service"
	"src/github.com/blackonyyx/cognizant/src/service/book"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)
type Data struct {
	BookService book.BookService
	BookController controller.BookController
	LoanService service.LoanService
	LoanController controller.LoanController
}
func NewData() *Data {
	var (
		bookService book.BookService = book.New()
		bookController controller.BookController = controller.NewBookController(bookService)
		loanService service.LoanService = service.New(bookService)
		loanController controller.LoanController = controller.NewLoanController(loanService)
	)
	return &Data{
		BookService: bookService,
		BookController: bookController,
		LoanService: loanService,
		LoanController: loanController,
	}
}
var DataLayer *Data


func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}
func SetupRouter() *gin.Engine {
	setupLogOutput()
	server := gin.New()
	DataLayer = NewData()
	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())
	// server.Use()
	server.POST("/add" , func (ctx *gin.Context) {

		res, err := DataLayer.BookController.Save(ctx)
		if err != nil {
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, res.String())
		}
	})
	server.GET("/books", func (ctx *gin.Context) {
		books := DataLayer.BookController.FindAll()
		fmt.Println(books)
		ctx.JSON(http.StatusOK, books)
	})

	server.GET("/read", func (ctx *gin.Context) { // /:key
		read, err := DataLayer.BookController.GetContent(ctx)
		if err != nil {
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, read)
		}

	})

	server.GET("/search", func(ctx *gin.Context) {
		books, err := DataLayer.BookController.GetBooks(ctx)
		if err != nil {
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, books)
		}
	})

	server.POST("/borrow", func(ctx *gin.Context) {
		resp, err := DataLayer.LoanController.BorrowBooks(ctx)
		if err != nil {
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, resp)
		}
	})

	server.GET("/loan", func(ctx *gin.Context) {
		resp, err := DataLayer.LoanController.GetLoanReceipt(ctx)
		if err != nil {
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, resp)
		}
	})

	server.POST("/return", func(ctx *gin.Context) {
		resp, err := DataLayer.LoanController.ReturnBooks(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, resp)
		}
	})
	server.POST("/extend", func(ctx *gin.Context) {
		resp, err := DataLayer.LoanController.ExtendLoan(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, resp)
		}
	})
	return server
}

