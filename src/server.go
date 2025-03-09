package src

import (
	"io"
	"log"
	"net/http"
	"os"
	"src/github.com/blackonyyx/cognizant/src/controller"
	"src/github.com/blackonyyx/cognizant/src/errormsg"
	"src/github.com/blackonyyx/cognizant/src/middlewares"
	"src/github.com/blackonyyx/cognizant/src/service"
	"src/github.com/blackonyyx/cognizant/src/service/book"
	mylog "src/github.com/blackonyyx/cognizant/src/log"
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


func setupLogOutput() *os.File {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	return f
}
func SetupRouter() *gin.Engine {
	f := setupLogOutput()
	log.SetOutput(f)
	server := gin.New()
	DataLayer = NewData()
	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())
	// server.Use()
	server.POST("/add" , func (ctx *gin.Context) {
		res, err := DataLayer.BookController.Save(ctx)
		if err != nil {
			mylog.CtxError(ctx, errormsg.ErrorMsgToStatusCode((err)), "Error in Adding Book")
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, res.String())
		}
	})
	server.GET("/books", func (ctx *gin.Context) {
		books := DataLayer.BookController.FindAll()
		ctx.JSON(http.StatusOK, books)
	})

	server.GET("/read", func (ctx *gin.Context) { // /:key
		read, err := DataLayer.BookController.GetContent(ctx)
		if err != nil {
			mylog.CtxError(ctx, errormsg.ErrorMsgToStatusCode((err)), "Error in Reading book")
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, read)
		}

	})

	server.GET("/search", func(ctx *gin.Context) {
		books, err := DataLayer.BookController.GetBooks(ctx)
		if err != nil {
			mylog.CtxError(ctx, errormsg.ErrorMsgToStatusCode((err)), "Error in Reading book")
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, books)
		}
	})

	server.POST("/borrow", func(ctx *gin.Context) {
		resp, err := DataLayer.LoanController.BorrowBooks(ctx)
		if err != nil {
			mylog.CtxError(ctx, errormsg.ErrorMsgToStatusCode((err)), "Error in borrowing book")
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, resp)
		}
	})

	server.GET("/loan", func(ctx *gin.Context) {
		resp, err := DataLayer.LoanController.GetLoanReceipt(ctx)
		if err != nil {
			mylog.CtxError(ctx, errormsg.ErrorMsgToStatusCode((err)), "Error in retrieving receipt")
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, resp)
		}
	})

	server.POST("/return", func(ctx *gin.Context) {
		resp, err := DataLayer.LoanController.ReturnBooks(ctx)
		if err != nil {
			mylog.CtxError(ctx, errormsg.ErrorMsgToStatusCode((err)), "Error in returning book")
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, resp)
		}
	})
	server.POST("/extend", func(ctx *gin.Context) {
		resp, err := DataLayer.LoanController.ExtendLoan(ctx)
		if err != nil {
			mylog.CtxError(ctx, errormsg.ErrorMsgToStatusCode(err), "Error in extending book loan")
			ctx.JSON(errormsg.ErrorMsgToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, resp)
		}
	})
	return server
}

