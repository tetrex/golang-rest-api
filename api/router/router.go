package router

import (
	"golang-rest-api/api/resource/book"
	"golang-rest-api/api/resource/health"

	_ "golang-rest-api/docs" // docs is generated by Swag CLI

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger" // middlerware for echo-swagger
	"gorm.io/gorm"
)

func New(db *gorm.DB, v *validator.Validate) *echo.Echo {
	r := echo.New()

	r.GET("/livez", health.Read)
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	g := r.Group("/v1")
	bookAPI := book.New(db, v)

	g.GET("/books", bookAPI.List)
	g.POST("/books", bookAPI.Create)
	g.GET("/books/:id", bookAPI.Read)
	g.PUT("/books/:id", bookAPI.Update)
	g.DELETE("/books/:id", bookAPI.Delete)

	return r
}
