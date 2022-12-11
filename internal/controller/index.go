package controller

import (
	"github.com/icemont/prime-numbers-tester/internal/request"
	"github.com/labstack/echo/v4"
	"math/big"
	"net/http"
)

type IndexController struct{}

func (ctrl *IndexController) GetIndex(c echo.Context) error {
	return c.HTML(http.StatusOK, `
			<h1>Prime Numbers Tester</h1>
			<h3>Usage Example:</h3>
			<pre>curl -X POST http://localhost:8888/ -F 'numbers=2' -F 'numbers=5'</pre>
		`)
}

func (ctrl *IndexController) PrimeNumberTester(c echo.Context) (err error) {
	r := new(request.NumbersRequest)
	if err = c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return err
	}

	result := make([]bool, len(r.Numbers))

	for i, num := range r.Numbers {
		result[i] = big.NewInt(num).ProbablyPrime(0)
	}

	return c.JSON(http.StatusOK, result)
}
