package main

import (
	"github.com/AndrewSDT/DnDSomeone/backend/internal/routing"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routing.LoadRoutes(e)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
