package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nmrshll/kata-fizzbuzz-api/services/fizzbuzzservice"
)

type FizzBuzzRoutes struct {
	fizzBuzzService fizzbuzzservice.FizzBuzzService
}

func (r *FizzBuzzRoutes) POSTFizzBuzz(c echo.Context) error {
	params := &fizzbuzzservice.POSTFizzBuzzParameters{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fizzBuzzOutput := r.fizzBuzzService.FizzBuzz(15, 3, 5, "fizz", "buzz")

	return c.JSON(http.StatusCreated, fizzBuzzOutput)
}
