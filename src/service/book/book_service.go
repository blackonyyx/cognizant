package book

import (
	"errors"
	"src/github.com/blackonyyx/cognizant/src/model"
)

type BookService interface {
	Save(model.Book) (model.Book, error)
	FindAll() []model.Book
	AddBook(string) (int64, error)
	ReturnBook(int64) (bool, error)
	BorrowBook(int64) (bool, error)
	GetContent(i int64) (model.BookContent, error)
}

type bookService struct {
	books        []model.Book
	bookContents map[int64]model.BookContent
}

// BorrowBook implements BookService.
func (service *bookService) BorrowBook(id int64) (bool, error) {
	if _, ok := service.bookContents[id]; !ok{
		return false, errors.New("book not found")
	}
	for _, i := range service.books {
		if i.Id == id {
			if (i.TotalStock > i.OnLoan) {
				i.OnLoan++
				return true, nil
			} else {
				return false, errors.New("book out of stock")
			}
		}
	}
	return false, errors.New("book not found")
}

// BorrowBook implements BookService.
func (service *bookService) ReturnBook(id int64) (bool, error) {
	if _, ok := service.bookContents[id]; !ok{
		return false, errors.New("book not found")
	}
	for _, i := range service.books {
		if i.Id == id {
			if (i.OnLoan > 0) {
				i.OnLoan--
				return true, nil
			} else {
				return false, errors.New("book return cannot be 0")
			}
		}
	}
	return false, errors.New("book not found")
}

func New() BookService {
	return &bookService{
		books:        []model.Book{},
		bookContents: map[int64]model.BookContent{},
	}
}

// FindAll implements BookService.
func (service *bookService) FindAll() []model.Book {
	return service.books
}

// Save implements BookService.
func (service *bookService) Save(book model.Book) (model.Book, error) {
	service.books = append(service.books, book)
	return book, nil
}

func (service *bookService) AddBook(bookContent string) (int64, error) {
	id := int64(len(service.bookContents)) + 1
	service.bookContents[id] = model.BookContent{Id: id, Content: bookContent}
	return id, nil
}

func (service *bookService) GetContent(i int64) (model.BookContent, error) {
	if i, ok := service.bookContents[i]; ok {
		return i, nil
	}
	return model.BookContent{}, errors.New("not found")
}
