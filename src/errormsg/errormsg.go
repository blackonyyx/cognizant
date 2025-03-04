package errormsg

import (
	"errors"
	"net/http"
)

var (
	INVALID_BINDING_INPUT = errors.New("invalid binding input")
	INVALID_INPUT = errors.New("invalid input")
	INVALID_STATUS = errors.New("invalid status")
	NOT_FOUND = errors.New("not found")
	OUT_OF_STOCK = errors.New("book out of stock")
	STOCK_ERROR = errors.New("impossible stock situation")
)
func ErrorMsgToStatusCode (err error) int {
	if err == NOT_FOUND {
		return http.StatusNoContent // In actuality use StatusNoContent
	} else if err == INVALID_STATUS || err == OUT_OF_STOCK || err == INVALID_BINDING_INPUT || err == INVALID_INPUT {
		return http.StatusBadRequest
	} else if err == STOCK_ERROR {
		return http.StatusInternalServerError
	} else {
		return http.StatusBadRequest
	}
}