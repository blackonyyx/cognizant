package reqbody

type LoanBooksRequest struct {
	BookIds []int64 `json:"book_ids" binding:"required"`
	Name string `json:"name"`
	Email string `json:"mail" validate:"required,email"`
}

type ReturnBooksRequest struct {
	ReceiptId int64 `json:"receipt_id"`
	LoanId []int64 `json:"loan_id"`
	Email string `json:"mail" binding:"required,email"`
}

type ExtensionRequest struct {
	ReceiptId int64 `json:"receipt_id" binding:"required"`
	LoanId []int64 `json:"loan_id"`
	Email string `json:"mail" binding:"required,email"`
}
