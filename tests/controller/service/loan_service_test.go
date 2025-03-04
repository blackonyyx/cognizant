package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"src/github.com/blackonyyx/cognizant/src"
	"src/github.com/blackonyyx/cognizant/src/errormsg"
	"src/github.com/blackonyyx/cognizant/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoan(t *testing.T) {
	t.Run("Loan Success", func(t *testing.T) {
		server := tests.SetupRouterWithBookData()
		w := httptest.NewRecorder()
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(tests.LOAN_REQUEST_1)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, tests.DOMAIN + tests.BORROW, &b)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		server.ServeHTTP(w, req)
		tmp := tests.LoanReceipt1()
		expected, _ := json.Marshal(tmp)
		assert.Equal(t, 200, w.Code, "API returns 200 for success")
		assert.Equal(t, string(expected), w.Body.String())
		assert.Equal(t, src.DataLayer.BookService.GetBook(1).OnLoan, int32(1), "stock changed successfully")


	})
	t.Run("Loan Failed Out of Stock", func(t *testing.T) {
		server := tests.SetupRouterWithBookAndLoanData()
		w := httptest.NewRecorder()
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(tests.LOAN_REQUEST_1)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, tests.DOMAIN + tests.BORROW, &b)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		server.ServeHTTP(w, req)
		tmp := tests.CreateErrorResp(errormsg.OUT_OF_STOCK)
		expected, _ := json.Marshal(tmp)
		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// some EOF error is added 
		assert.Contains(t, w.Body.String(), string(expected))
	})
	// will omit field binding valuation tests for subsequent api
	t.Run("Loan Email Invalid", func(t *testing.T) {
		server := tests.SetupRouterWithBookAndLoanData()
		w := httptest.NewRecorder()
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(tests.INVALID_EMAIL_LOAN_REQUEST)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, tests.DOMAIN + tests.BORROW, &b)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		server.ServeHTTP(w, req)
		tmp := "Field validation for 'Email' failed on the 'email' tag"
		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), tmp)
		
	})
}

func TestExtension(t *testing.T) {
	t.Run("Extend Loan", func(t *testing.T) {
		server := tests.SetupRouterWithBookAndLoanData()
		w := httptest.NewRecorder()
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(tests.LOAN_EXTENSION)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, tests.DOMAIN + tests.EXTEND, &b)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		server.ServeHTTP(w, req)
		tmp := tests.ExtendReceipt1()
		expected, _ := json.Marshal(tmp)
		assert.Equal(t, 200, w.Code, "API returns 200 for success")
		assert.Equal(t, string(expected), w.Body.String())
	})
	t.Run("Extend Loan Twice Failed", func(t *testing.T) {
		server := tests.SetupRouterWithBookAndLoanData()
		a := httptest.NewRecorder()
		
		w := httptest.NewRecorder()
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(tests.LOAN_EXTENSION)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, tests.DOMAIN + tests.EXTEND, &b)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		server.ServeHTTP(a, req)
		err = json.NewEncoder(&b).Encode(tests.LOAN_EXTENSION)
		if err != nil {
			t.Fatal(err)
		}
		req = httptest.NewRequest(http.MethodPost, tests.DOMAIN + tests.EXTEND, &b)
		// tmp := tests.ExtendReceipt1()
		tmp := tests.CreateErrorResp(errormsg.INVALID_STATUS)
		expected, _ := json.Marshal(tmp)

		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, string(expected), w.Body.String())
	})
	t.Run("Invalid Loan Id", func(t *testing.T) {
		server := tests.SetupRouterWithBookAndLoanData()
		w := httptest.NewRecorder()
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(tests.INVALID_LOAN_ID_LOAN_EXTENSION)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, tests.DOMAIN + tests.EXTEND, &b)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		tmp := tests.CreateErrorResp(errormsg.NOT_FOUND)
		expected, _ := json.Marshal(tmp)

		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code, "Loan Id not found")
		assert.Equal(t, string(expected), w.Body.String())
	})
}

func TestReturnBook(t *testing.T) {
	t.Run("Return successfully", func(t *testing.T) {
		server := tests.SetupRouterWithBookAndLoanData()
		w := httptest.NewRecorder()
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(tests.LOAN_RETURN)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, tests.DOMAIN + tests.RETURN, &b)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		server.ServeHTTP(w, req)
		tmp := tests.ReturnReceipt1()
		expected, _ := json.Marshal(tmp)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expected), w.Body.String())
		assert.Equal(t, src.DataLayer.BookService.GetBook(1).OnLoan, int32(0), "stock changed successfully")
	})
	t.Run("Invalid Loan Id", func(t *testing.T) {
		server := tests.SetupRouterWithBookAndLoanData()
		w := httptest.NewRecorder()
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(tests.INVALID_ID_LOAN_RETURN)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, tests.DOMAIN + tests.RETURN, &b)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		tmp := tests.CreateErrorResp(errormsg.NOT_FOUND)
		expected, _ := json.Marshal(tmp)

		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code, "Loan Id not found")
		assert.Equal(t, string(expected), w.Body.String())
	})
}