package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/icemont/prime-numbers-tester/internal/controller"
	"github.com/icemont/prime-numbers-tester/internal/request"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.Validator = &request.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	indexController := new(controller.IndexController)

	e.GET("/", indexController.GetIndex)
	e.POST("/", indexController.PrimeNumberTester)

	return e
}
