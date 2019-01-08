package api

import (
	// "encoding/json"
	"fmt"

	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nmrshll/kata-fizzbuzz-api/services/fizzbuzzservice"
)

const PORT = 8080

func StartServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// FizzBuzz route
	fizzBuzzRoutes := FizzBuzzRoutes{fizzBuzzService: fizzbuzzservice.FizzBuzzService{}}
	e.POST("/fizzbuzz", fizzBuzzRoutes.POSTFizzBuzz)

	// Start server
	fmt.Printf("Listening on port %s...\n", strconv.Itoa(PORT))
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(PORT)))
}
