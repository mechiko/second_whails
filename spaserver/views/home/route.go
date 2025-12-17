package home

import (
	"korrectkm/domain"
	"korrectkm/domain/models/modeltrueclient"
	"korrectkm/guiconnect"
	"korrectkm/reductor"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *page) Routes() error {
	base := "/" + t.modelType.String()
	t.Echo().GET(base, t.Index)
	t.Echo().GET(base+"/reset", t.Reset)
	t.Echo().GET(base+"/license", t.license)
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

func (t *page) Reset(c echo.Context) error {
	return nil
}

func (t *page) license(c echo.Context) error {
	tcModel, err := reductor.Model[*modeltrueclient.TrueClientModel](domain.TrueClient, t)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := guiconnect.StartDialog(t, tcModel); err != nil {
		return t.ServerError(c, err)
	}
	if err := reductor.SetModel(tcModel, false); err != nil {
		return t.ServerError(c, err)
	}
	return c.NoContent(204)
}
