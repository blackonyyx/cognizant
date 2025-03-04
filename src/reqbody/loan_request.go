package reqbody

type LoanBooksRequest struct {
	BookIds []int64 `json:"book_ids" binding:"required"`
	Name string `json:"name"`
	Email string `json:"email" binding:"required,email"`
}

type ReturnBooksRequest struct {
	LoanIds []int64 `json:"loan_ids"`
	Name string `json:"name"`
	Email string `json:"email" binding:"required,email"`
}

type ExtensionRequest struct {
	LoanIds []int64 `json:"loan_ids"`
	Email string `json:"email" binding:"required,email"`
	Name string `json:"name"`

}
