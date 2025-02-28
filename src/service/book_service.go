package service

import "src/github.com/blackonyyx/cognizant/src/model"

type BookService interface {
	Save(model.Book) model.Book
	FindAll() []model.Book
	// Borro
}

type bookService struct {
	books []model.Book
}

func New() BookService {
	return &bookService{}
}

// FindAll implements BookService.
func (service *bookService) FindAll() []model.Book {
	panic("unimplemented")
}

// Save implements BookService.
func (service *bookService) Save(book model.Book) model.Book {
	service.books = append(service.books, book)
	return book
}
