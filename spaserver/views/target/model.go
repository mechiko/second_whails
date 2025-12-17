package target

import (
	"fmt"
	"korrectkm/domain"
	"korrectkm/reductor"
	"time"
)

type TargetModel struct {
	model      domain.Model
	Title      string
	Gtin       string
	Pg         string
	CisStatus  map[string][]domain.TargetCis
	Filter     domain.TargetFilter
	TargetCis  []domain.TargetCis
	Updated    time.Time
	Progress   int   // прогресс опроса
	IsProgress bool  // true если идет процесс загрузки для отображения прогресса
	Error      error // массив ошибок
}

var _ domain.Modeler = (*TargetModel)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func NewModel(app domain.Apper) (*TargetModel, error) {
	model := &TargetModel{
		model:     domain.Target,
		Title:     "Выборка по фильтру",
		TargetCis: make([]domain.TargetCis, 0),
		Updated:   time.Time{},
		Filter: domain.TargetFilter{
			Pagination: domain.FilterPagination{
				PerPage:          1000,
				LastEmissionDate: "2025-12-10T22:00:00.000Z",
				Direction:        0,
				Sgtin:            "000000000000000000000",
			},
		},
	}
	if err := model.ReadState(app); err != nil {
		return nil, fmt.Errorf("model %v read state %w", model.model, err)
	}
	return model, nil
}

// инициализируем модель вида
func (t *page) InitData(app domain.Apper) (interface{}, error) {
	model, err := NewModel(app)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	err = reductor.SetModel(model, false)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (m *TargetModel) SyncToStore(_ domain.Apper) (err error) {
	return err
}

// читаем состояние приложения
func (m *TargetModel) ReadState(app domain.Apper) (err error) {
	return nil
}

func (m *TargetModel) Copy() (interface{}, error) {
	// shallow copy that`s why fields is simple
	dst := *m
	return &dst, nil
}

func (a *TargetModel) Model() domain.Model {
	return a.model
}

func (a *TargetModel) Save(_ domain.Apper) (err error) {
	return nil
}

// всегда возвращает true означает проверки нет всегда ок
func (m *TargetModel) License(_ domain.Apper) bool {
	return true
}

func (m *TargetModel) Pgs() map[string]string {
	mm := map[string]string{}
	for k, v := range domain.ProductGroupByAlias {
		mm[k] = v.Name
	}
	return mm
}
