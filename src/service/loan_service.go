package service

import (
	"errors"
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/src/service/book"
)

type LoanService interface {
	CreateLoan([]model.Book) (model.LoanReceipt, error)
	UpdateStatus([]model.BookLoan) (model.LoanReceipt, error)
	GetLoanReceipt(int64, string) (model.LoanReceipt, error)
}

type loanService struct {
	BookService  *book.BookService
	Loans        []model.BookLoan
	LoanReceipts []model.LoanReceipt
}

// CreateLoan implements LoanService.
func (l *loanService) CreateLoan([]model.Book) (model.LoanReceipt, error) {
	panic("unimplemented")
}

// GetLoanReceipt implements LoanService.
func (l *loanService) GetLoanReceipt(receiptId int64, email string) (model.LoanReceipt, error) {
	for _, i := range l.LoanReceipts {
		if i.Id == receiptId {
			return i, nil
		} else if i.Email == email {
			return i, nil
		}
	}
	return model.LoanReceipt{}, errors.New("not found")
}

// UpdateStatus implements LoanService.
func (l *loanService) UpdateStatus([]model.BookLoan) (model.LoanReceipt, error) {
	panic("unimplemented")
}

func New(books *book.BookService) LoanService {
	return &loanService{
		BookService:  books,
		Loans:        []model.BookLoan{},
		LoanReceipts: []model.LoanReceipt{},
	}
}
