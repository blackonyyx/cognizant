package service

import "src/github.com/blackonyyx/cognizant/src/model"

type BookService interface {
	Save(model.Book) model.Book
	FindAll() []model.Book
	UpdateStock(int64, int32) model.Book
}

type bookService struct {
	books []model.Book
}

// UpdateStock implements BookService.
func (service *bookService) UpdateStock(bookId int64, stockChange int32) model.Book {
	for _, book := range service.books {
		if book.Id == bookId {
			book.OnLoan = book.OnLoan + stockChange
			return book
		}
	}
	return model.Book{}
}

func New() BookService {
	return &bookService{}
}

// FindAll implements BookService.
func (service *bookService) FindAll() []model.Book {
	return service.books
}

// Save implements BookService.
func (service *bookService) Save(book model.Book) model.Book {
	service.books = append(service.books, book)
	return book
}
