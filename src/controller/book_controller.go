package controller

import (
	"errors"
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/src/reqbody"
	"src/github.com/blackonyyx/cognizant/src/service/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookController interface {
	FindAll() []model.Book
	Save(ctx *gin.Context) (model.Book, error)
	GetContent(ctx *gin.Context) (model.BookContent, error)
}

var validate *validator.Validate
type controller struct {
	service book.BookService

}

func New(service book.BookService) BookController {
	validate = validator.New()
	validate.RegisterValidation("my_validator", nil)
	return &controller{
		service: service,
	}
}

// FindAll implements BookController.
func (c *controller) FindAll() []model.Book {
	return c.service.FindAll()
}

// Save implements BookController.
func (c *controller) Save(ctx *gin.Context) (model.Book, error) {
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

func (c *controller) GetContent(ctx *gin.Context) (model.BookContent, error) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return model.BookContent{}, errors.New("Invalid Id input")
	}
	// todo add key to access borrowed book.
	book, err := c.service.GetContent(int64(idInt))
	if err != nil {
		return model.BookContent{}, err
	}
	return book, nil
}


