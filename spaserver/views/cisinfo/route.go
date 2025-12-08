package cisinfo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (t *page) Routes() error {
	// Serve static and media files under /static/ and /uploads/ path.
	base := "/" + t.modelType.String()
	t.Echo().GET(base, t.Index)
	t.Echo().GET(base+"/reset", t.Reset)
	t.Echo().POST(base+"/info", t.Info)
	return nil
}

func (t *page) Index(c echo.Context) error {
	data, err := t.PageData()
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("index", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

// сброс к начальному состоянию
func (t *page) Reset(c echo.Context) error {
	data, err := t.InitData(t)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("index", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) Info(c echo.Context) error {
	cis := c.FormValue("cis")
	// Обрезаем служебный разделитель GS1 (FNC1, 0x1D), если он есть.
	if idx := strings.IndexByte(cis, '\x1D'); idx >= 0 {
		cis = cis[:idx]
	}

	if cis == "" {
		return t.ServerError(c, fmt.Errorf("пустой код"))
	}
	if len(cis) < 16 {
		return t.ServerError(c, fmt.Errorf("мало знаков для кода"))
	}
	if err := t.info(cis); err != nil {
		return t.ServerError(c, err)
	}
	data, err := t.PageData()
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("info", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}
