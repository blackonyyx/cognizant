package tests

import (
	"src/github.com/blackonyyx/cognizant/src"
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/src/reqbody"
	"src/github.com/blackonyyx/cognizant/src/service"

	"time"
)


var (
	DOMAIN = "http://localhost:3000/"
	BORROW = "borrow"
	EXTEND = "extend"
	RETURN = "return"
	READ_PARAMS1 = "read?id=1"
	INVALID_READ_PARAMS = "read?id=100"
	
	INVALID_SEARCH_PARAMS1 = "search?id=10"
	INVALID_SEARCH_PARAMS2 = "search?title=asjdnakjsd"
	INVALID_SEARCH_PARAMS3 = "search?title=ipsum&author=john&id=1"
	SEARCH_PARAMS1 = "search?id=1"
	SEARCH_PARAMS2 = "search?title=ipsum&author=john"
	SEARCH_PARAMS3 = "search?title=ipsum"

	INVALID_LOAN_PARAMS1 = "loan?id=-1"
	LOAN_PARAMS1 = "loan?id=1"


)

var (
	BOOK_CONTENT1 = model.BookContent{
		Id: 1,
		Content: "The quick brown fox jumps over the wall",
	}
	BOOK_CONTENT2 = model.BookContent{
		Id: 2,
		Content: "Askadasdasb",
	}
	BOOK_CONTENT3 = model.BookContent{
		Id: 3,
		Content : "Formula 1 Racing123123",
	}
	BOOK1  = model.Book{
		Id: 1,
		Title: "Alice In Wonderland",
		Author: "Lewis Carol",
		Description: "Alice In wonderland Is A girl",
		TotalStock: 1,
		OnLoan: 0,
	}
	BOOK2 = model.Book{
		Id: 2,
        Title: "lorem ipsum",
        Author: "john doe",
        Description: "lasdasdasdaasd",
        TotalStock: 12,
        OnLoan: 0,
	}
	BOOK3 = model.Book {
		Id: 3,
		Title: "ipsum World",
		Author: "john Alexander Hamilton",
		Description: "",
		TotalStock: 2,
		OnLoan: 0,
	}
	ADD_REQUEST1 = reqbody.SaveBookRequest{
		Title : "Alice In Wonderland",
		Author: "Lewis Carol",
		Description: "Alice In wonderland Is A girl",
		TotalStock: 1,
		Content : "The quick brown fox jumps over the wall",
	}
	ADD_REQUEST2 = reqbody.SaveBookRequest{
		Title : "lorem ipsum",
		Author: "john doe",
		Description: "lasdasdasdaasd",
		TotalStock: 12,
		Content : "Askadasdasb",
	}
	ADD_REQUEST3 = reqbody.SaveBookRequest{
		Title: "ipsum World",
		Author: "john Alexander Hamilton",
		Description: "",
		TotalStock: 2,
		Content : "Formula 1 Racing123123",
	}
	LOAN_REQUEST_1 = reqbody.LoanBooksRequest {
		BookIds: []int64{1, 3},
		Email: "stephen@m.com",
	}
	LOAN_REQUEST_2 = reqbody.LoanBooksRequest {
		BookIds: []int64{2, 3},
		Email: "john@u.com",
		Name: "John",
	}
	INVALID_EMAIL_LOAN_REQUEST = reqbody.LoanBooksRequest {
		BookIds: []int64{2, 3},
		Email: "johnu.com",
		Name: "John",
	}
	
	LOAN_EXTENSION = reqbody.ExtensionRequest {
		LoanIds: []int64{1, 2},
		Email: "john@u.com",
		Name: "John",
	}
	INVALID_LOAN_ID_LOAN_EXTENSION = reqbody.ExtensionRequest {
		LoanIds: []int64{1, 0},
		Email: "john@u.com",
		Name: "John",
	}
	INVALID_EMAIL_LOAN_EXTENSION = reqbody.ExtensionRequest {
		LoanIds: []int64{1, 2},
		Email: "johnu.com",
		Name: "John",
	}
	
	LOAN_RETURN = reqbody.ReturnBooksRequest {
		LoanIds: []int64{1, 2},
		Email: "john@u.com",
		Name: "John",
	}

	INVALID_ID_LOAN_RETURN = reqbody.ReturnBooksRequest {
		LoanIds: []int64{1, 9},
		Email: "john@u.com",
		Name: "John",
	}

	INVALID_EMAIL_LOAN_RETURN = reqbody.ReturnBooksRequest {
		LoanIds: []int64{1, 2},
		Email: "john@com",
		Name: "John",
	}
)

func CreateErrorResp(err error) map[string]interface{} {
	v := map[string]interface{}{
		"error": err.Error(),
	}
	return v
}

func LoanReceipt1() model.LoanReceipt {
	now := time.Now()
	next := now.AddDate(0, 0, 28)
	loans := []model.BookLoan{
		model.BookLoan{
			LoanId: 1,
			ReceiptId: 1,
			BookId: 1,
			StartDate: now.Unix(),
			EndDate: next.Unix(),
			Status: service.ON_LOAN,
		},
		model.BookLoan{
			LoanId: 2,
			ReceiptId: 1,
			BookId: 3,
			StartDate: now.Unix(),
			EndDate: next.Unix(),
			Status: service.ON_LOAN,
		},
	}
	return model.LoanReceipt{
		Id: 1,
		Email: "stephen@m.com",
		Loans: loans,
		StartDate: now.Unix(),
		EndDate: next.Unix(),
	}
}

func LoanReceipt2() model.LoanReceipt {
	now := time.Now()
	next := now.AddDate(0, 0, 28)
	loans := []model.BookLoan{
		model.BookLoan{
			LoanId: 3,
			ReceiptId: 2,
			BookId: 2,
			StartDate: now.Unix(),
			EndDate: next.Unix(),
			Status: service.ON_LOAN,
		},
		model.BookLoan{
			LoanId: 4,
			ReceiptId: 2,
			BookId: 3,
			StartDate: now.Unix(),
			EndDate: next.Unix(),
			Status: service.ON_LOAN,
		},
	}
	return model.LoanReceipt{
		Id: 1,
		Email: "stephen@m.com",
		Loans: loans,
		StartDate: now.Unix(),
		EndDate: next.Unix(),
	}
}

func ExtendReceipt1() model.LoanReceipt {
	now, _ := src.DataLayer.LoanService.GetLoanReceipt(1)
	t := now.StartDate
	next := time.Unix(t, 0).AddDate(0, 0, 28).AddDate(0, 0, 21)
	loans := []model.BookLoan{
		model.BookLoan{
			LoanId: 1,
			ReceiptId: 2,
			BookId: 1,
			StartDate: t,
			EndDate: next.Unix(),
			Status: service.EXTENDED,
		},
		model.BookLoan{
			LoanId: 2,
			ReceiptId: 2,
			BookId: 3,
			StartDate: t,
			EndDate: next.Unix(),
			Status: service.EXTENDED,
		},
	}
	return model.LoanReceipt{
		Id: 2,
		Name: "John",
		Email: "john@u.com",
		Loans: loans,
	}
}

func ReturnReceipt1() model.LoanReceipt {
	now, _ := src.DataLayer.LoanService.GetLoanReceipt(1)
	loans := now.Loans
	for i := range loans {
		loans[i].Status = service.RETURNED
	}
	return model.LoanReceipt{
		Id: 2,
		Name: "John",
		Email: "john@u.com",
		Loans: loans,
	}
}
