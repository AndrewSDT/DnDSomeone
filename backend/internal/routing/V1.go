package routing

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo) {
	g := e.Group("/v1")
	g.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})
}
