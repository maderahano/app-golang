package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooksController(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		expectStatus int
	}{
		{
			name:         "get book normal",
			path:         "/books",
			expectStatus: http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &books)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}

func TestGetBookController(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		expectStatus int
		// expectBodyStartWith string
	}{
		{
			name:         "get book normal",
			path:         "/books",
			expectStatus: http.StatusOK,
			// expectBodyStartWith: "{\"status\":\"success\",\"books\":[",
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &books)

			if err != nil {
				assert.Error(t, err, "error")
			}
			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestCreateBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get book normal",
			path:       "/books/6",
			expectCode: http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &books)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}

func TestUpdateBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get book normal",
			path:       "/books/6",
			expectCode: http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &books)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}

func TestDeleteBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get book normal",
			path:       "/books/6",
			expectCode: http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &books)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}
