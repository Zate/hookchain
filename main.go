package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func root(c echo.Context) error {
	return c.String(http.StatusOK, "hookchain")
}

func hook(c echo.Context) error {
	return c.String(http.StatusOK, "hook handler")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))
	e.Use(middleware.Secure())

	e.GET("/", root)
	e.POST("/hook", hook)
	e.Logger.Fatal(e.Start(":3000"))
}
