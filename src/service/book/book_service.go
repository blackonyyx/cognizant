package book

import (
	"errors"
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/src/reqbody"
	"strings"

	"github.com/samber/lo"
)

type BookService interface {
	Save(model.Book) (model.Book, error)
	FindAll() []model.Book
	AddBook(string) (int64, error)
	ReturnBooks([]int64) (bool, error)
	BorrowBooks([]int64) (bool, error)
	GetContent(i int64) (model.BookContent, error)
	FindBooks(reqbody.FindBookRequest) ([]model.Book, error)
}

type bookService struct {
	books        []model.Book
	bookContents map[int64]model.BookContent
}

// FindBook implements BookService.
func (service *bookService) FindBooks(req reqbody.FindBookRequest) ([]model.Book, error) {
	if _, exist := service.bookContents[req.BookId] ;req.BookId != 0 &&  exist {
		book, found := lo.Find(service.books, func(b model.Book) bool {
			return b.Id == req.BookId
		})
		if !found {
			return nil,  errors.New("not found")
		}
		return []model.Book{book}, nil
	}
	books := lo.Filter(service.books, func(b model.Book, _ int) bool {
		test := true
		if len(req.Author) > 2 {
			test = test && (strings.Contains(b.Author, req.Author))
		}
		if len(req.Title) > 2 {
			test = test && (strings.Contains(b.Title, req.Title))
		}
		return test
	})
	if len(books) == 0 {
		return books, errors.New("not found")
	}
	return books, nil
}

// BorrowBook implements BookService.
func (service *bookService) BorrowBooks(id []int64) (bool, error) {
	m := lo.SliceToMap(id, func(f int64) (int64, bool) {
		return f, true
	})
	var list []*model.Book
	for i, _ := range service.books {
		if m[service.books[i].Id] {
			list = append(list, &service.books[i])
		}
	}
	if len(id) != len(list) {
		return false, errors.New("some book id does not exist in the library")
	}

	for _, ptr := range list {
		if (*ptr).TotalStock > (*ptr).OnLoan {
			(*ptr).OnLoan++
		} else {
			return false, errors.New("some book is out of stock")
		}
	}
	return true, nil
}

// BorrowBook implements BookService.
func (service *bookService) ReturnBooks(id []int64) (bool, error) {
	m := lo.SliceToMap(id, func(f int64) (int64, bool) {
		return f, true
	})
	var list []*model.Book
	for i, _ := range service.books {
		if m[service.books[i].Id] {
			list = append(list, &service.books[i])
		}
	}
	if len(id) != len(list) {
		return false, errors.New("some book id mentioned does not exist in the library")
	}

	for _, ptr := range list {
		if (*ptr).TotalStock > (*ptr).OnLoan {
			(*ptr).OnLoan++
		} else {
			return false, errors.New("some book is out of stock")
		}
		if (*ptr).OnLoan > 0 {
			(*ptr).OnLoan--
			return true, nil
		} else {
			return false, errors.New("book return cannot be 0")
		}
	}
	return true, nil
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
