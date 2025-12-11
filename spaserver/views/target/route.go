package target

import (
	"errors"
	"korrectkm/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *page) Routes() error {
	// Serve static and media files under /static/ and /uploads/ path.
	base := "/" + t.modelType.String()
	t.Echo().GET(base, t.Index)
	t.Echo().GET(base+"/reset", t.Reset)
	t.Echo().POST(base+"/scan", t.Scan)
	t.Echo().GET(base+"/progress", t.progress)
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

func (t *page) Scan(c echo.Context) error {
	gtin := c.FormValue("gtin")
	data, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	data.Filter.Filter.Gtins = []string{gtin}
	data.Filter.Filter.ProductGroups = []string{"beer"}
	data.Filter.Filter.States = []domain.FilterState{
		{Status: "INTRODUCED"},
		{Status: "APPLIED"},
	}
	data.Filter.Pagination.PerPage = 1000
	if data.IsProgress {
		// уже запущена операция
		// return c.NoContent(204)
		return t.ServerError(c, errors.New("уже запущена операция, дождитесь ее окончания"))
	}
	// запускаем прогресс чтобы отобразить на странице он отображается когда больше 0
	data.IsProgress = true
	if err := t.ModelUpdate(data); err != nil {
		return t.ServerError(c, err)
	}
	go func() {
		if err := t.scan(); err != nil {
			data, _ := t.PageModel()
			data.Error = err
			t.ModelUpdate(data)
		}
	}()
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) progress(c echo.Context) error {
	data, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	if data.IsProgress {
		if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("progress", data)); err != nil {
			return t.ServerError(c, err)
		}
	} else {
		if data.Error != nil {
			return t.ServerError(c, err)
		} else {
			if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page_stop", data)); err != nil {
				return t.ServerError(c, err)
			}
		}
	}
	return nil
}
