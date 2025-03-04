package controller

import (
	"src/github.com/blackonyyx/cognizant/src/errormsg"
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/src/reqbody"
	"src/github.com/blackonyyx/cognizant/src/service/book"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	FindAll() []model.Book
	Save(ctx *gin.Context) (model.Book, error)
	GetContent(ctx *gin.Context) (model.BookContent, error)
	GetBooks(ctx *gin.Context) ([]model.Book, error)
}

type bookController struct {
	service book.BookService

}

func NewBookController(service book.BookService) BookController {
	return &bookController{
		service: service,
	}
}

// FindAll implements BookController.
func (c *bookController) FindAll() []model.Book {
	return c.service.FindAll()
}

// Save implements BookController.
func (c *bookController) Save(ctx *gin.Context) (model.Book, error) {
	var req reqbody.SaveBookRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return model.Book{}, errormsg.INVALID_INPUT
	}
	book , err := c.service.SaveBook(req)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (c *bookController) GetBooks(ctx *gin.Context) ([]model.Book, error) {
	var req reqbody.FindBookRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		return nil, errormsg.INVALID_BINDING_INPUT
	}
	books, err := c.service.FindBooks(req)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (c *bookController) GetContent(ctx *gin.Context) (model.BookContent, error) {
	id := ctx.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return model.BookContent{}, errormsg.INVALID_INPUT
	}
	// todo add key to access borrowed book.
	book, err := c.service.GetContent(int64(idInt))
	if err != nil {
		return model.BookContent{}, err
	}
	return book, nil
}


