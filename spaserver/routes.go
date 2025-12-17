package spaserver

import (
	"fmt"
	"korrectkm/spaserver/views/adjust"
	"korrectkm/spaserver/views/cisinfo"
	"korrectkm/spaserver/views/footer"
	"korrectkm/spaserver/views/gtin"
	"korrectkm/spaserver/views/header"
	"korrectkm/spaserver/views/home"
	"korrectkm/spaserver/views/innfias"
	"korrectkm/spaserver/views/kmstate"
	"korrectkm/spaserver/views/menu"
	"korrectkm/spaserver/views/money"
	"korrectkm/spaserver/views/root"
	"korrectkm/spaserver/views/target"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.design/x/clipboard"
)

// маршрутизация приложения
func (s *Server) Routes() http.Handler {
	s.loadViews()
	s.server.GET("/", s.Index)
	s.server.GET("/copy/:value", s.copy) // переход/загрузка на текущую страницу
	s.server.GET("/page", s.Page)        // переход/загрузка на текущую страницу
	s.server.GET("/page/reset", s.Reset) // переход/загрузка на текущую страницу
	s.server.GET("/sse", s.Sse)
	return s.server
}

// вызывает рендеринг по имени страницы вида в s.activePage
func (s *Server) Page(c echo.Context) error {
	ap := s.activePage
	if _, ok := s.views[ap]; !ok {
		return s.ServerError(c, fmt.Errorf("not found active page %s", ap))
	}
	view := s.views[ap]
	return view.Index(c)
}

// вызывает рендеринг по имени страницы вида в s.activePage
func (s *Server) Reset(c echo.Context) error {
	ap := s.activePage
	if _, ok := s.views[ap]; !ok {
		return s.ServerError(c, fmt.Errorf("not found active page %s", ap))
	}
	view := s.views[ap]
	return view.Reset(c)
}

// загружаем все виды
func (s *Server) loadViews() {
	// view header
	view1 := header.New(s)
	s.views[view1.ModelType()] = view1
	view1.Routes()
	// view footer
	view2 := footer.New(s)
	s.views[view2.ModelType()] = view2
	view2.Routes()
	view2.InitData(s)
	// view home
	view3 := home.New(s)
	s.views[view3.ModelType()] = view3
	view3.Routes()
	view3.InitData(s)
	// kmstate
	view4 := kmstate.New(s)
	s.views[view4.ModelType()] = view4
	view4.Routes()
	view4.InitData(s)
	// view root
	view5 := root.New(s)
	s.views[view5.ModelType()] = view5
	// menu
	view6 := menu.New(s)
	s.views[view6.ModelType()] = view6
	view6.Routes()
	// innfias
	view7 := innfias.New(s)
	s.views[view7.ModelType()] = view7
	view7.Routes()
	view7.InitData(s)
	// money
	view8 := money.New(s)
	s.views[view8.ModelType()] = view8
	view8.Routes()
	view8.InitData(s)
	// CisInfo
	view9 := cisinfo.New(s)
	s.views[view9.ModelType()] = view9
	view9.Routes()
	view9.InitData(s)
	// Adjust
	view10 := adjust.New(s)
	s.views[view10.ModelType()] = view10
	view10.Routes()
	view10.InitData(s)
	// Target
	view11 := target.New(s)
	s.views[view11.ModelType()] = view11
	view11.Routes()
	view11.InitData(s)
	// Gtin
	view12 := gtin.New(s)
	s.views[view12.ModelType()] = view12
	view12.Routes()
	view12.InitData(s)
	// header инициализируем последним нужны все виды сервера в списке
	view1.InitData(s)
	view6.InitData(s)
}

func (s *Server) Index(c echo.Context) error {
	if err := c.Render(http.StatusOK, "root", map[string]interface{}{
		"template": "index",
		"data":     &struct{}{},
	}); err != nil {
		return s.ServerError(c, err)
	}
	return nil
}

func (s *Server) copy(c echo.Context) error {
	value := c.Param("value")
	clipboard.Write(clipboard.FmtText, []byte(value))
	// walk.Clipboard().SetText(value)
	// s.SetFlush("инфо", fmt.Sprintf("значение %s скопировано", value))
	s.SetFlush("скопировано в буфер значение "+value, "info")
	return c.String(200, "")
}
