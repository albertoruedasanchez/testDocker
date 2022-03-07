package main

import (
	"github.com/arueda/testDocker/services"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	services.Service(e)
}
