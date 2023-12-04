package book

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type API struct{}

func (a *API) List(c echo.Context) error {
	return c.String(http.StatusOK, "list")
}

func (a *API) Create(c echo.Context) error {
	return c.String(http.StatusOK, "create")
}

func (a *API) Read(c echo.Context) error {
	return c.String(http.StatusOK, "read")
}

func (a *API) Update(c echo.Context) error {
	return c.String(http.StatusOK, "update")
}

func (a *API) Delete(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}
