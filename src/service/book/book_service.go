package book

import (
	"fmt"
	"src/github.com/blackonyyx/cognizant/src/errormsg"
	"src/github.com/blackonyyx/cognizant/src/log"
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/src/reqbody"
	"strings"

	"github.com/samber/lo"
)

type BookService interface {
	SaveBook(reqbody.SaveBookRequest) (model.Book, error)
	FindAll() []model.Book
	ReturnBooks([]int64) (bool, error)
	BorrowBooks([]int64) (bool, error)
	GetContent(i int64) (model.BookContent, error)
	
	FindBooks(reqbody.FindBookRequest) ([]model.Book, error)
	// testing method
	GetBook(i int64) (model.Book)
}

type bookService struct {
	books        []model.Book
	bookContents map[int64]model.BookContent
}
func (service *bookService) GetBook(i int64) (model.Book) {
	book, _ := lo.Find(service.books, func(b model.Book) bool {
		return b.Id == i
	})
	return book
}

// SaveBook implements BookService.
func (service *bookService) SaveBook(req reqbody.SaveBookRequest) (model.Book, error) {
	var book int64
	if req.Id == 0 {
		tmp, err := service.AddBook(req.Content)
		book = tmp
		if err != nil {
			log.Warning("[SaveBook]: ", err.Error())
			return model.Book{}, err
		}
	} else {
		book = req.Id
		if _, err := service.GetContent(book) ; err != nil {
			return model.Book{}, err
		}
	}
	var bookStock model.Book
	bookStock.Id = book
	bookStock.Title = req.Title
	bookStock.Author = req.Author
	bookStock.Description = req.Description
	bookStock.TotalStock = req.TotalStock
	log.Info("[SaveBook]:", fmt.Sprintf("%#v", bookStock))

	return service.Save(bookStock)
}

// FindBook implements BookService.
func (service *bookService) FindBooks(req reqbody.FindBookRequest) ([]model.Book, error) {
	if req.BookId != 0 && (req.Author != "" || req.Title != "") {
		log.Warning(fmt.Sprintf("[FindBooks] Req: %#v, Invalid Request", req), errormsg.INVALID_INPUT.Error())
		return nil, errormsg.INVALID_INPUT
	}
	if _, exist := service.bookContents[req.BookId]; req.BookId != 0 && exist {
		book, found := lo.Find(service.books, func(b model.Book) bool {
			return b.Id == req.BookId
		})
		if !found {
			log.Warning(fmt.Sprintf("[FindBooks] Req: %#v, Book not found", req))

			return nil, errormsg.NOT_FOUND
		}
		return []model.Book{book}, nil
	} else if req.BookId != 0 && !exist {
		log.Warning(fmt.Sprintf("[FindBooks] Req: %#v, Book not found", req))
		return nil, errormsg.NOT_FOUND
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
		log.Warning(fmt.Sprintf("[FindBooks] %#v, Book not found", req))
		return books, errormsg.NOT_FOUND
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
		log.Warning(fmt.Sprintf("[BorrowBook] %#v, Book not found", id))
		return false, errormsg.NOT_FOUND
	}

	for _, ptr := range list {
		if (*ptr).TotalStock > (*ptr).OnLoan {
			(*ptr).OnLoan++
		} else {
			log.Warning(fmt.Sprintf("[BorrowBook] %#v, Book out of stock", id))
			return false, errormsg.OUT_OF_STOCK
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
			log.Warning(fmt.Sprintf("[ReturnBooks]: %v, some id not found", id))
			return false, errormsg.NOT_FOUND
	}

	for _, ptr := range list {
		if (*ptr).OnLoan > 0 {
			(*ptr).OnLoan--
			return true, nil
		} else {
			log.Warning("[ReturnBooks]: ", errormsg.STOCK_ERROR.Error())
			return false, errormsg.STOCK_ERROR
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
	log.Info(fmt.Sprintf("[FindAll] Books %#v", service.books))
	return service.books
}

// Save implements BookService.
func (service *bookService) Save(book model.Book) (model.Book, error) {
	service.books = append(service.books, book)
	log.Info(fmt.Sprintf("[Save] Book Entry Added, Contents: %#v", book))
	return book, nil
}

func (service *bookService) AddBook(bookContent string) (int64, error) {
	id := int64(len(service.bookContents)) + 1
	service.bookContents[id] = model.BookContent{Id: id, Content: bookContent}
	log.Info(fmt.Sprintf("[AddBook] Book Added, Contents: %#v", service.bookContents[id]))
	return id, nil
}

func (service *bookService) GetContent(i int64) (model.BookContent, error) {
	if i, ok := service.bookContents[i]; ok {
		log.Info(fmt.Sprintf("[GetContent] Book found, Contents: %#v", i))
		return i, nil
	}
	log.Warning("[GetContent] Book Id not found")
	return model.BookContent{}, errormsg.NOT_FOUND
}
