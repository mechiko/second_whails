package kmstate

import (
	"fmt"
	"korrectkm/reductor"
	"korrectkm/repo"
	"net/http"

	"github.com/labstack/echo/v4"
)

type loaderConfig struct {
	formField   string
	errorMsg    string
	updateModel func(*KmStateModel, string) error
	fetchCodes  func(interface{}) ([]string, error)
}

func (t *page) genericLoad(c echo.Context, cfg loaderConfig) error {
	value := c.FormValue(cfg.formField)
	if value == "" {
		return t.ServerError(c, fmt.Errorf(cfg.errorMsg))
	}

	rp, err := repo.GetRepository()
	if err != nil {
		return t.ServerError(c, err)
	}

	zn, err := rp.LockZnak()
	if err != nil {
		return t.ServerError(c, err)
	}
	defer rp.UnlockZnak(zn)

	model, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}

	if err := cfg.updateModel(model, value); err != nil {
		return t.ServerError(c, err)
	}

	model.CisIn = nil
	inCis, err := cfg.fetchCodes(value)
	if err != nil {
		return t.ServerError(c, err)
	}

	model.CisIn = append(model.CisIn, inCis...)
	if len(model.CisIn) == 0 {
		return t.ServerError(c, fmt.Errorf("получено 0 КМ"))
	}

	if err := reductor.SetModel(model, false); err != nil {
		return t.ServerError(c, err)
	}

	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page", model)); err != nil {
		return t.ServerError(c, err)
	}

	return nil
}
