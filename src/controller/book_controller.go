package controller

import (
	"errors"
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
	var book int64
	if req.Id == 0 {
		book, err = c.service.AddBook(req.Content)
		if err != nil {
			return model.Book{}, err
		}
	} else {
		book = req.Id
		if _, err := c.service.GetContent(book) ; err != nil {
			return model.Book{}, err
		}
	}
	var bookStock model.Book
	bookStock.Id = book
	bookStock.Title = req.Title
	bookStock.Author = req.Author
	bookStock.Description = req.Description
	bookStock.TotalStock = req.TotalStock

	if err != nil {
		// err handle
		return model.Book{}, err
	}
	return c.service.Save(bookStock)
}

func (c *bookController) GetBooks(ctx *gin.Context) ([]model.Book, error) {
	var req reqbody.FindBookRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		return nil, errors.New("invalid input")
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
		return model.BookContent{}, errors.New("invalid id input")
	}
	// todo add key to access borrowed book.
	book, err := c.service.GetContent(int64(idInt))
	if err != nil {
		return model.BookContent{}, err
	}
	return book, nil
}


