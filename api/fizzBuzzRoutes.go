package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nmrshll/kata-fizzbuzz-api/services/fizzbuzzservice"
)

// FizzBuzzRoutes holds all fizzbuzz-routes related methods in one dependency
type FizzBuzzRoutes struct {
	fizzBuzzService fizzbuzzservice.FizzBuzzService
}

// POSTFizzBuzz responds to POST /fizzbuzz
// expects a body with all 5 parameters (limit,int1,int2,str1,str2) (see fixzzbuzzservice.POSTFizzBuzzParameters)
func (r *FizzBuzzRoutes) POSTFizzBuzz(c echo.Context) error {
	params := &fizzbuzzservice.POSTFizzBuzzParameters{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fizzBuzzOutput := r.fizzBuzzService.FizzBuzz(15, 3, 5, "fizz", "buzz")

	return c.JSON(http.StatusOK, fizzBuzzOutput)
}
