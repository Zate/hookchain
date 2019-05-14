package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func Test_root(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, root(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "hookchain")
	}
}

func Test_hook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, hook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "hook handler")
	}
}

func Test_main(t *testing.T) {
	e := echo.New()
	// req := httptest.NewRequest(http.MethodGet, "/", nil)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)

	// if assert.NoError(t, hook(c)) {
	// 	assert.Equal(t, http.StatusOK, rec.Code)
	// 	assert.Contains(t, rec.Body.String(), "hook handler")
	// }
	e.Use(middleware.Logger())
	assert.NotNil(t, middleware.Logger())
	e.Use(middleware.Recover())
	assert.NotNil(t, middleware.CORS())
	// e.HideBanner = true
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.POST},
	// }))
	// e.Use(middleware.Secure())

	// e.GET("/", root)
	// e.POST("/hook", hook)
	// e.Logger.Fatal(e.Start(":3000"))
}
