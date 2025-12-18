package domain

import (
	"github.com/labstack/echo/v4"
)

type IServer interface {
	Apper
	Echo() *echo.Echo
	ServerError(c echo.Context, err error) error
	SetActivePage(Model)
	SetFlush(string, string)
	RenderString(name string, data interface{}) (str string, err error)
	Reload()
}
