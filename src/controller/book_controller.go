package controller

import (
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/src/service"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	FindAll() []model.Book
	Save(ctx *gin.Context) model.Book
}

type controller struct {
	service service.BookService
}

func New(service service.BookService) BookController {
	return controller{
		service: service,
	}
}

// FindAll implements BookController.
func (c *controller) FindAll() []model.Book {
	return c.service.FindAll()
}

// Save implements BookController.
func (c *controller) Save(ctx *gin.Context) model.Book {
	var book model.Book
	ctx.BindJSON(&book)
	return c.service.Save(book)
}


