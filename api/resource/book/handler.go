package book

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type API struct{}

// List godoc
//
//	@summary		List books
//	@description	List books
//	@tags			books
//	@accept			json
//	@produce		json
//	@success		200	{array}		DTO
//	@failure		500	{object}	err.Error
//	@router			/books [get]
func (a *API) List(c echo.Context) error {
	return c.String(http.StatusOK, "list")
}

// Create godoc
//
//	@summary		Create book
//	@description	Create book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			body	body	Form	true	"Book form"
//	@success		201
//	@failure		400	{object}	err.Error
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/books [post]
func (a *API) Create(c echo.Context) error {
	return c.String(http.StatusOK, "create")
}

// Read godoc
//
//	@summary		Read book
//	@description	Read book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			id	path		string	true	"Book ID"
//	@success		200	{object}	DTO
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		500	{object}	err.Error
//	@router			/books/{id} [get]
func (a *API) Read(c echo.Context) error {
	return c.String(http.StatusOK, "read")
}

// Update godoc
//
//	@summary		Update book
//	@description	Update book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			id		path	string	true	"Book ID"
//	@param			body	body	Form	true	"Book form"
//	@success		200
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/books/{id} [put]
func (a *API) Update(c echo.Context) error {
	return c.String(http.StatusOK, "update")
}

// Delete godoc
//
//	@summary		Delete book
//	@description	Delete book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			id	path	string	true	"Book ID"
//	@success		200
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		500	{object}	err.Error
//	@router			/books/{id} [delete]
func (a *API) Delete(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}
