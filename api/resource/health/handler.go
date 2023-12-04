package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Read(c echo.Context) error {
	return c.String(http.StatusOK, "health")
}
