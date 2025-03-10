package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"src/github.com/blackonyyx/cognizant/src/errormsg"
	"src/github.com/blackonyyx/cognizant/src/model"
	"src/github.com/blackonyyx/cognizant/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// router := src.SetupRouter()
	code := m.Run()
	os.Exit(code)
}

func TestReadBookSuccess(t *testing.T) {
	server := tests.SetupRouterWithBookData()
	t.Run("valid BookId", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tests.DOMAIN + tests.READ_PARAMS1, nil)
		tmp := tests.BOOK_CONTENT1
		expected, _ := json.Marshal(tmp)
		server.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code, "API returns 200 for success")
		assert.Equal(t, string(expected), w.Body.String())
	})
	t.Run("Invalid BookId", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tests.DOMAIN + tests.INVALID_READ_PARAMS, nil)
		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code, "API returns No Content Found")
		assert.Equal(t, "", w.Body.String())
	})
}

func TestSearchBook(t *testing.T) {
	server := tests.SetupRouterWithBookData()
	t.Run("valid exact search", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tests.DOMAIN + tests.SEARCH_PARAMS1, nil)
		tmp := []model.Book{tests.BOOK1}
		expected, _ := json.Marshal(tmp)
		server.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code, "API returns 200 for success")
		assert.Equal(t, string(expected), w.Body.String())
	})
	t.Run("invalid exact search id", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tests.DOMAIN + tests.INVALID_SEARCH_PARAMS1, nil)
		
		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code,  "API returns No Content Found")
		assert.Equal(t, "", w.Body.String())
	})
	t.Run("invalid partial search id", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tests.DOMAIN + tests.INVALID_SEARCH_PARAMS2, nil)
		
		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code,  "API returns No Content Found")
		assert.Equal(t, "", w.Body.String())
	})
	t.Run("valid partial multi search", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tests.DOMAIN + tests.SEARCH_PARAMS2, nil)
		tmp := []model.Book{tests.BOOK2, tests.BOOK3}
		expected, _ := json.Marshal(tmp)
		server.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code, "API returns 200 for success")
		assert.Equal(t, string(expected), w.Body.String())
	})
	t.Run("valid partial single field search", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tests.DOMAIN + tests.SEARCH_PARAMS3, nil)
		tmp := []model.Book{tests.BOOK2, tests.BOOK3}
		expected, _ := json.Marshal(tmp)
		server.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code, "API returns 200 for success")
		assert.Equal(t, string(expected), w.Body.String())
	})
	t.Run("invalid all field search", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, tests.DOMAIN + tests.INVALID_SEARCH_PARAMS3, nil)
		tmp := tests.CreateErrorResp(errormsg.INVALID_INPUT)
		expected, _ := json.Marshal(tmp)
		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, string(expected), w.Body.String())
	})
}
