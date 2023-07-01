package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// BaseRouter implements Router.
func (r router) BaseRouter() {
	r.route.GET("/health_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1Previx := r.route.Group("/api")
	{
		v1Previx.GET("/health_check", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
	}
}
