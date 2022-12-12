package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/icemont/prime-numbers-tester/internal/request"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPrimeNumberTesterMethod(t *testing.T) {
	var testCases = []struct {
		name       string
		whenQuery  string
		expectBody string
		expectCode int
	}{
		{
			name:       "ok1",
			whenQuery:  "numbers=2&numbers=3",
			expectBody: "[true,true]\n",
			expectCode: http.StatusOK,
		},
		{
			name:       "ok2",
			whenQuery:  "numbers=1&numbers=5",
			expectBody: "[false,true]\n",
			expectCode: http.StatusOK,
		},
		{
			name:      "nok1",
			whenQuery: "test=error",
			expectBody: "{\"message\":\"Key: 'NumbersRequest.Numbers' Error:Field validation for 'Numbers' failed " +
				"on the 'gt' tag\"}\n",
			expectCode: http.StatusBadRequest,
		},
		{
			name:      "nok2",
			whenQuery: "numbers=error&numbers=5",
			expectBody: "{\"message\":\"code=400, message=strconv.ParseInt: parsing \\\"error\\\": invalid syntax, " +
				"internal=strconv.ParseInt: parsing \\\"error\\\": invalid syntax\"}\n",
			expectCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = &request.CustomValidator{Validator: validator.New()}
			e.Use(middleware.Recover())
			e.Use(middleware.Logger())

			indexController := new(IndexController)
			e.POST("/", indexController.PrimeNumberTester)

			req := httptest.NewRequest(http.MethodPost, "http://localhost/", strings.NewReader(tc.whenQuery))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectBody, rec.Body.String())
			assert.Equal(t, tc.expectCode, rec.Code)
		})
	}
}
