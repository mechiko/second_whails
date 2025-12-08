package header

import (
	"fmt"
	"korrectkm/domain"
)

type HeaderModel struct {
	IApp
	Title string
	Items MenuItemSlice
	model domain.Model
}

type MenuItem struct {
	Name   string
	Title  string
	Active bool
	Desc   string
	Svg    string
}

type MenuItemSlice []*MenuItem

var _ domain.Modeler = (*HeaderModel)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func NewModel(app IApp) (*HeaderModel, error) {
	model := &HeaderModel{
		IApp:  app,
		model: domain.Header,
		Title: "Меню",
		Items: make(MenuItemSlice, 0),
	}
	if err := model.ReadState(app); err != nil {
		return nil, fmt.Errorf("model MenuModel read state %w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (m *HeaderModel) SyncToStore(app domain.Apper) (err error) {
	return err
}

// читаем состояние приложения
func (m *HeaderModel) ReadState(app domain.Apper) (err error) {
	return nil
}

func (m *HeaderModel) Copy() (interface{}, error) {
	// shallow copy that`s why fields is simple
	dst := *m
	return &dst, nil
}

func (a *HeaderModel) Model() domain.Model {
	return a.model
}

func (a *HeaderModel) Save(_ domain.Apper) (err error) {
	return nil
}

func (a *HeaderModel) SyncActive(active string) error {
	for _, m := range a.Items {
		m.Active = false
		if m.Name == active {
			m.Active = true
		}
	}
	return nil
}

func (a *HeaderModel) ActiveTitle() string {
	for _, m := range a.Items {
		if m.Active {
			return m.Desc
		}
	}
	return ""
}

// всегда возвращает true означает проверки нет всегда ок
func (m *HeaderModel) License() bool {
	return true
}
