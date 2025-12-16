package target

import (
	"errors"
	"fmt"
	"korrectkm/domain"
	"korrectkm/reductor"
	"korrectkm/spaserver/views/kmstate"
	"korrectkm/spaserver/views/menu"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (t *page) Routes() error {
	// Serve static and media files under /static/ and /uploads/ path.
	base := "/" + t.modelType.String()
	t.Echo().GET(base, t.Index)
	t.Echo().GET(base+"/reset", t.Reset)
	t.Echo().POST(base+"/scan", t.Scan)
	t.Echo().GET(base+"/progress", t.progress)
	t.Echo().GET(base+"/tostatus", t.toStatus)
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
	pg := c.FormValue("pg")
	data, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	data.Gtin = gtin
	data.Pg = pg
	data.Filter.Filter.Gtins = []string{data.Gtin}
	data.Filter.Filter.ProductGroups = []string{data.Pg}
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
			data.IsProgress = false
			data.Updated = time.Time{}
			t.ModelUpdate(data)
			return
		}
		data, _ := t.PageModel()
		data.Updated = time.Now()
		data.Error = nil
		data.IsProgress = false
		t.ModelUpdate(data)
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
			t.ServerError(c, data.Error)
			data.Error = nil
			t.ModelUpdate(data)
			return nil
		} else {
			if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page_stop", data)); err != nil {
				return t.ServerError(c, err)
			}
		}
	}
	return nil
}

func (t *page) toStatus(c echo.Context) error {
	modelDst, err := kmstate.NewModel(t)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	data, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	modelDst.File = "отобрано по фильтру"
	for _, cis := range data.TargetCis {
		modelDst.CisIn = append(modelDst.CisIn, cis.Cis)
	}
	err = reductor.SetModel(modelDst, false)
	if err != nil {
		return fmt.Errorf("target page model update %w", err)
	}
	modelMenu, err := reductor.Model[*menu.MenuModel](domain.Menu)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	modelMenu.SyncActive(string(domain.KMState))
	err = reductor.SetModel(modelMenu, false)
	if err != nil {
		return fmt.Errorf("target page model update %w", err)
	}
	t.SetActivePage(domain.KMState)
	t.Reload()
	return nil
}
