package model

type LoanReceipt struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email" validate:"required,email"`
	Loans []BookLoan `json:"loan_ids"`
	StartDate int64 `json:"start_date"` // for convenience
	EndDate int64 `json:"end_date"`
}

type BookLoan struct {
	ReceiptId int64 `json:"receipt_id"`
	LoanId int64 `json:"loan_id"`
	BookId int64 `json:"book_id"`
	StartDate int64 `json:"start_date"`
	EndDate int64 `json:"end_date"`
	Status int64 `json:"status"` // status 1: on loan, 2: returned, 3: extended, 4: expired
}