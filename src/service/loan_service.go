package service

import "src/github.com/blackonyyx/cognizant/src/model"

type LoanService interface {
	CreateLoan([]model.BookLoan) model.LoanReceipt
	UpdateStatus([]model.BookLoan) model.LoanReceipt
	GetLoanReceipt(int64, string) model.LoanReceipt
}