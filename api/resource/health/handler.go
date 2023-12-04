package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Read godoc
//
//	@summary		Read health
//	@description	Read health
//	@tags			health
//	@success		200
//	@router			/../livez [get]
func Read(c echo.Context) error {
	return c.String(http.StatusOK, "health")
}
