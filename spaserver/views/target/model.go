package target

import (
	"fmt"
	"korrectkm/domain"
	"korrectkm/reductor"
	"time"
)

type TargetModel struct {
	model   domain.Model
	Title   string
	Balance *domain.Balance
	Updated time.Time
}

var _ domain.Modeler = (*TargetModel)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func NewModel(app domain.Apper) (*TargetModel, error) {
	model := &TargetModel{
		model:   domain.Target,
		Title:   "Информация по ИНН",
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
func (m *TargetModel) License() bool {
	return true
}
