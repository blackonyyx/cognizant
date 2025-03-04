package tests

import (
	"src/github.com/blackonyyx/cognizant/src"

	"github.com/gin-gonic/gin"
)

func SetupEmptyRouter() *gin.Engine {
	router := src.SetupRouter()
	return router
}

func SetupRouterWithBookData() *gin.Engine {
	router := SetupEmptyRouter()
	src.DataLayer.BookService.SaveBook(ADD_REQUEST1)
	src.DataLayer.BookService.SaveBook(ADD_REQUEST2)
	src.DataLayer.BookService.SaveBook(ADD_REQUEST3)
	return router
}

func SetupRouterWithBookAndLoanData() *gin.Engine {
	router := SetupEmptyRouter()
	src.DataLayer.BookService.SaveBook(ADD_REQUEST1)
	src.DataLayer.BookService.SaveBook(ADD_REQUEST2)
	src.DataLayer.BookService.SaveBook(ADD_REQUEST3)
	src.DataLayer.LoanService.CreateLoan(LOAN_REQUEST_1)
	return router
}