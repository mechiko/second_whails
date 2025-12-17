package adjust

import (
	"fmt"
	"korrectkm/domain"
	"korrectkm/reductor"
	"time"
)

type AdjustModel struct {
	model   domain.Model
	Title   string
	Balance *domain.Balance
	Updated time.Time
}

var _ domain.Modeler = (*AdjustModel)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func NewModel(app domain.Apper) (*AdjustModel, error) {
	model := &AdjustModel{
		model:   domain.Adjust,
		Title:   "Корректировка",
		Balance: &domain.Balance{},
		Updated: time.Time{},
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
func (m *AdjustModel) SyncToStore(_ domain.Apper) (err error) {
	return err
}

// читаем состояние приложения
func (m *AdjustModel) ReadState(app domain.Apper) (err error) {
	return nil
}

func (m *AdjustModel) Copy() (interface{}, error) {
	// shallow copy that`s why fields is simple
	dst := *m
	return &dst, nil
}

func (a *AdjustModel) Model() domain.Model {
	return a.model
}

func (a *AdjustModel) Save(_ domain.Apper) (err error) {
	return nil
}

// всегда возвращает true означает проверки нет всегда ок
func (m *AdjustModel) License(_ domain.Apper) bool {
	return true
}
