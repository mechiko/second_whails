package kmstate

import (
	"errors"
	"fmt"
	"korrectkm/reductor"
	"korrectkm/repo"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mechiko/utility"
)

func (t *page) Routes() error {
	// Serve static and media files under /static/ and /uploads/ path.
	base := "/" + t.modelType.String()
	t.Echo().GET(base, t.Index)
	t.Echo().GET(base+"/selectfile", t.selectFile)
	t.Echo().GET(base+"/reset", t.Reset)
	t.Echo().GET(base+"/progress", t.progress)
	t.Echo().GET(base+"/search", t.search)
	t.Echo().GET(base+"/excel/:status/:statusex", t.excel)
	t.Echo().POST(base+"/orderload", t.orderLoad)
	t.Echo().POST(base+"/utilload", t.utilLoad)
	t.Echo().POST(base+"/atkload", t.atkLoad)
	t.Echo().POST(base+"/gtinload", t.gtinLoad)
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

func (t *page) selectFile(c echo.Context) error {
	model, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	model.File, err = utility.DialogOpenFile([]utility.FileType{utility.Csv}, "", t.Pwd())
	if err != nil {
		return t.ServerError(c, err)
	}
	err = model.loadCisFromFile()
	if err != nil {
		return t.ServerError(c, err)
	}
	err = reductor.SetModel(model, false)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page", model)); err != nil {
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
		if len(data.Errors) > 0 {
			err := strings.Join(data.Errors, "<BR>")
			data.Errors = make([]string, 0)
			t.ModelUpdate(data)
			return t.ServerError(c, errors.New(err))
		} else {
			if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page_stop", data)); err != nil {
				return t.ServerError(c, err)
			}
		}
	}
	return nil
}

func (t *page) search(c echo.Context) error {
	data, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
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
	go t.Search()
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("process_cis", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) excel(c echo.Context) error {
	status := c.Param("status")
	statusEx := c.Param("statusex")
	file, err := utility.DialogSaveFile(utility.Excel, "", ".")
	if err != nil {
		return t.ServerError(c, err)
	}
	file = fmt.Sprintf("%s_%s", file, status)
	data, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	allcis := data.CisOut
	arrStatus := make([]string, 0)
	for _, cis := range allcis {
		if (cis.Status == status) && (cis.StatusEx == statusEx) {
			arrStatus = append(arrStatus, cis.Cis)
		}
	}
	size := data.ExcelChunkSize
	if size <= 0 {
		size = 30000
	}
	fNames, err := t.ToExcel(arrStatus, file, size)
	if err != nil {
		return t.ServerError(c, err)
	}
	msg := fmt.Sprintf("записано %d КМ в %d файл(а,ов)", len(arrStatus), len(fNames))
	t.SetFlush(msg, "info")
	// возвращаем имя файла для отображения на форме
	// out := fmt.Sprintf("записано %d кодов маркировки %d файлов", len(arrStatus), len(fNames))
	// c.String(200, out)
	c.NoContent(204)
	return nil
}

func (t *page) orderLoad(c echo.Context) error {
	order := c.FormValue("order")
	orderId, err := strconv.Atoi(order)
	if err != nil {
		t.Logger().Errorf("order not number %w", err)
		return t.ServerError(c, fmt.Errorf("введите число больше 0 ид заказа"))
	}
	rp, err := repo.GetRepository()
	if err != nil {
		return t.ServerError(c, err)
	}
	zn, err := rp.LockZnak()
	if err != nil {
		return t.ServerError(c, err)
	}
	defer func() {
		rp.UnlockZnak(zn)
	}()
	model, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	model.OrderId = orderId
	model.CisIn = nil
	inCis, err := zn.OrderCodes(int64(orderId))
	if err != nil {
		return t.ServerError(c, err)
	}
	model.CisIn = append(model.CisIn, inCis...)
	if len(model.CisIn) == 0 {
		return t.ServerError(c, fmt.Errorf("получено 0 КМ"))
	}
	err = reductor.SetModel(model, false)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page", model)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) utilLoad(c echo.Context) error {
	util := c.FormValue("util")
	utilId, err := strconv.Atoi(util)
	if err != nil {
		t.Logger().Errorf("utilisation report not number %w", err)
		return t.ServerError(c, fmt.Errorf("введите число больше 0 ид отчета нанесения"))
	}
	rp, err := repo.GetRepository()
	if err != nil {
		return t.ServerError(c, err)
	}
	zn, err := rp.LockZnak()
	if err != nil {
		return t.ServerError(c, err)
	}
	defer func() {
		rp.UnlockZnak(zn)
	}()
	model, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	model.UtilisationId = utilId
	model.CisIn = nil
	inCis, err := zn.UtilisationCodes(int64(utilId))
	if err != nil {
		return t.ServerError(c, err)
	}
	model.CisIn = append(model.CisIn, inCis...)
	if len(model.CisIn) == 0 {
		return t.ServerError(c, fmt.Errorf("получено 0 КМ"))
	}
	err = reductor.SetModel(model, false)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page", model)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) atkLoad(c echo.Context) error {
	atk := c.FormValue("atk")
	atkId, err := strconv.Atoi(atk)
	if err != nil {
		t.Logger().Errorf("atk not number %w", err)
		return t.ServerError(c, fmt.Errorf("введите число больше 0 ид атк"))
	}
	rp, err := repo.GetRepository()
	if err != nil {
		return t.ServerError(c, err)
	}
	zn, err := rp.LockZnak()
	if err != nil {
		return t.ServerError(c, err)
	}
	defer func() {
		rp.UnlockZnak(zn)
	}()
	model, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	model.AtkId = atkId
	model.CisIn = nil
	inCis, err := zn.AtkCodes(int64(atkId))
	if err != nil {
		return t.ServerError(c, err)
	}
	model.CisIn = append(model.CisIn, inCis...)
	if len(model.CisIn) == 0 {
		return t.ServerError(c, fmt.Errorf("получено 0 КМ"))
	}
	err = reductor.SetModel(model, false)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page", model)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) gtinLoad(c echo.Context) error {
	gtin := c.FormValue("gtin")
	if len(gtin) == 0 {
		return t.ServerError(c, fmt.Errorf("введите GTIN"))
	}
	rp, err := repo.GetRepository()
	if err != nil {
		return t.ServerError(c, err)
	}
	zn, err := rp.LockZnak()
	if err != nil {
		return t.ServerError(c, err)
	}
	defer func() {
		rp.UnlockZnak(zn)
	}()
	model, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	model.Gtin = gtin
	model.CisIn = nil
	inCis, err := zn.GtinCodes(gtin)
	if err != nil {
		return t.ServerError(c, err)
	}
	model.CisIn = append(model.CisIn, inCis...)
	if len(model.CisIn) == 0 {
		return t.ServerError(c, fmt.Errorf("получено 0 КМ"))
	}
	err = reductor.SetModel(model, false)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("page", model)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}
