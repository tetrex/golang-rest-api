package router

import (
	"golang-rest-api/api/resource/book"
	"golang-rest-api/api/resource/health"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	r := echo.New()

	r.GET("/livez", health.Read)

	g := r.Group("/admin")
	bookAPI := &book.API{}

	g.GET("/books", bookAPI.List)
	g.POST("/books", bookAPI.Create)
	g.GET("/books/{id}", bookAPI.Read)
	g.PUT("/books/{id}", bookAPI.Update)
	g.DELETE("/books/{id}", bookAPI.Delete)

	return r
}
