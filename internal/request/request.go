package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	NumbersRequest struct {
		Numbers []int64 `form:"numbers" validate:"gt=0,dive,required,number,min=1"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
