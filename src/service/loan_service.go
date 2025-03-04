package service

import (
	"src/github.com/blackonyyx/cognizant/src/errormsg"
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/src/reqbody"
	"src/github.com/blackonyyx/cognizant/src/service/book"
	"time"
)

const (
	ON_LOAN  = iota
	RETURNED = iota
	EXTENDED = iota
	EXPIRED  = iota
)

type LoanService interface {
	CreateLoan(reqbody.LoanBooksRequest) (model.LoanReceipt, error)
	ExtendStatus(reqbody.ExtensionRequest) (model.LoanReceipt, error)
	ReturnLoan(reqbody.ReturnBooksRequest) (model.LoanReceipt, error)
	GetLoanReceipt(int64) (model.LoanReceipt, error)
}

type loanService struct {
	BookService  book.BookService
	Loans        map[int]model.BookLoan
	LoanReceipts []model.LoanReceipt
}

// ReturnLoan implements LoanService.
func (l *loanService) ReturnLoan(req reqbody.ReturnBooksRequest) (model.LoanReceipt, error) {
	var bookIds []int64
	for _, loanId := range req.LoanIds {
		loan, ok := l.Loans[int(loanId)]
		if !ok {
			return model.LoanReceipt{}, errormsg.NOT_FOUND
		}
		if loan.Status == RETURNED {
			return model.LoanReceipt{}, errormsg.INVALID_STATUS
		}
		bookIds = append(bookIds, loan.BookId)

	}
	
	if ok, err := l.BookService.ReturnBooks(bookIds); !ok {
		return model.LoanReceipt{}, err
	}
	var loans []model.BookLoan
	for _, id := range req.LoanIds {
		tmp := l.Loans[int(id)]
		tmp.Status = RETURNED
		loans = append(loans, tmp)
		l.Loans[int(id)] = tmp
	}
	receipt := model.LoanReceipt{
		Id:    int64(len(l.LoanReceipts) + 1),
		Name:  req.Name,
		Email: req.Email,
		Loans: loans,
	}
	l.LoanReceipts = append(l.LoanReceipts, receipt)
	return receipt, nil
}

// CreateLoan implements LoanService.
func (l *loanService) CreateLoan(userData reqbody.LoanBooksRequest) (model.LoanReceipt, error) {
	if borrowed, err := l.BookService.BorrowBooks(userData.BookIds); !borrowed {
		return model.LoanReceipt{}, err
	}
	receiptId := len(l.LoanReceipts) + 1
	loanTime := time.Now()
	returnTime := loanTime.AddDate(0, 0, 28)
	var loans []model.BookLoan
	for _, i := range userData.BookIds {
		loanId := int64(len(l.Loans)) + 1
		loan := model.BookLoan{
			ReceiptId: int64(receiptId),
			LoanId:    loanId,
			BookId:    i,
			StartDate: loanTime.Unix(),
			EndDate:   returnTime.Unix(),
			Status:    int64(ON_LOAN),
		}
		loans = append(loans, loan)
		l.Loans[int(loanId)] = loan
	}
	res := model.LoanReceipt{
		Id:        int64(receiptId),
		Name:      userData.Name,
		Email:     userData.Email,
		Loans:     loans,
		StartDate: loanTime.Unix(),
		EndDate:   returnTime.Unix(),
	}
	l.LoanReceipts = append(l.LoanReceipts, res)
	// userData.BookIds
	return res, nil
}

// GetLoanReceipt implements LoanService.
func (l *loanService) GetLoanReceipt(receiptId int64) (model.LoanReceipt, error) {
	for _, i := range l.LoanReceipts {
		if i.Id == receiptId {
			return i, nil
		}
	}
	return model.LoanReceipt{}, errormsg.NOT_FOUND
}

// UpdateStatus implements LoanService.
func (l *loanService) ExtendStatus(req reqbody.ExtensionRequest) (model.LoanReceipt, error) {
	now := time.Now()
	for _, loanId := range req.LoanIds {
		loan, ok := l.Loans[int(loanId)]
		if !ok {
			return model.LoanReceipt{}, errormsg.NOT_FOUND
		}
		endDate := time.Unix(loan.EndDate, 0)
		if now.After(endDate) {
			return model.LoanReceipt{}, errormsg.INVALID_STATUS
		}
		if loan.Status == RETURNED {
			return model.LoanReceipt{}, errormsg.INVALID_STATUS
		}
		if loan.Status == EXTENDED {
			return model.LoanReceipt{}, errormsg.INVALID_STATUS
		}
	}
	var loans []model.BookLoan
	receiptId := int64(len(l.LoanReceipts) + 1)
	for _, id := range req.LoanIds {
		tmp := l.Loans[int(id)]
		tmp.EndDate = time.Unix(l.Loans[int(id)].EndDate, 0).AddDate(0, 0, 21).Unix()
		tmp.ReceiptId = receiptId
		tmp.Status = EXTENDED
		loans = append(loans, tmp)
		l.Loans[int(id)] = tmp
	}
	receipt := model.LoanReceipt{
		Id:    receiptId,
		Name:  req.Name,
		Email: req.Email,
		Loans: loans,
	}
	l.LoanReceipts = append(l.LoanReceipts, receipt)
	return receipt, nil
}

func New(books book.BookService) LoanService {
	return &loanService{
		BookService:  books,
		Loans:        map[int]model.BookLoan{},
		LoanReceipts: []model.LoanReceipt{},
	}
}
