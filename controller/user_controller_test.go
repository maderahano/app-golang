package controller

import (
	"app-golang/config"
	"app-golang/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	config.InitDB()
	e := echo.New()

	return e
}

func TestGetUsersController(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		expectStatus int
		// expectBodyStartWith string
	}{
		{
			name:         "get user normal",
			path:         "/users",
			expectStatus: http.StatusOK,
			// expectBodyStartWith: "{\"status\":\"success\",\"users\":[",
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &users)

			if err != nil {
				assert.Error(t, err, "error")
			}
			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		expectStatus int
		// expectBodyStartWith string
	}{
		{
			name:         "get user normal",
			path:         "/users",
			expectStatus: http.StatusOK,
			// expectBodyStartWith: "{\"status\":\"success\",\"users\":[",
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &users)

			if err != nil {
				assert.Error(t, err, "error")
			}
			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestCreateUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get user normal",
			path:       "/users/21",
			expectCode: http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &users)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}

func TestUpdateUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get user normal",
			path:       "/users/21",
			expectCode: http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &users)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}

func TestDeleteUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get user normal",
			path:       "/users/21",
			expectCode: http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			err := json.Unmarshal([]byte(body), &users)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}

func TestLoginUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, LoginUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var user []model.UserResponse
			err := json.Unmarshal([]byte(body), &user)

			if err != nil {
				assert.Error(t, err, "error")
			}
		}
	}
}
