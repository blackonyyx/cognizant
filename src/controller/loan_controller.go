package controller

import (
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/src/reqbody"
	"src/github.com/blackonyyx/cognizant/src/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanController interface {
	GetLoanReceipt(ctx *gin.Context) (model.LoanReceipt, error)
	BorrowBooks(ctx *gin.Context) (model.LoanReceipt, error)
	ReturnBooks(ctx *gin.Context) (model.LoanReceipt, error)
	ExtendLoan(ctx *gin.Context) (model.LoanReceipt, error)
}

type loanController struct {
	service service.LoanService
}

// BorrowBooks implements LoanController.
func (l *loanController) BorrowBooks(ctx *gin.Context) (model.LoanReceipt, error) {
	var req reqbody.LoanBooksRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return model.LoanReceipt{}, err
	}
	receipt , err := l.service.CreateLoan(req)
	if err != nil {
		return model.LoanReceipt{}, err
	}
	return receipt, nil
}

// ExtendLoan implements LoanController.
func (l *loanController) ExtendLoan(ctx *gin.Context) (model.LoanReceipt, error) {
	var req reqbody.ExtensionRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return model.LoanReceipt{}, err
	}
	receipt , err := l.service.ExtendStatus(req)
	if err != nil {
		return model.LoanReceipt{}, err
	}
	return receipt, nil
}

// GetLoanReceipt implements LoanController.
func (l *loanController) GetLoanReceipt(ctx *gin.Context) (model.LoanReceipt, error) {
	receiptId := ctx.Query("id")
	idInt, err := strconv.Atoi(receiptId)
	if err != nil {
		return model.LoanReceipt{}, err
	}
	receipt, err := l.service.GetLoanReceipt(int64(idInt))
	if err != nil {
		return model.LoanReceipt{}, err
	}
	return receipt, nil
}

// ReturnBooks implements LoanController.
func (l *loanController) ReturnBooks(ctx *gin.Context) (model.LoanReceipt, error) {
	var req reqbody.ReturnBooksRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return model.LoanReceipt{}, err
	}
	receipt , err := l.service.ReturnLoan(req)
	if err != nil {
		return model.LoanReceipt{}, err
	}
	return receipt, nil
}

func NewLoanController(service service.LoanService) LoanController {
	return &loanController{
		service: service,
	}
}
