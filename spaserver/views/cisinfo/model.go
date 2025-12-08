package cisinfo

import (
	"fmt"
	"korrectkm/domain"
	"korrectkm/reductor"
	"time"
)

type CisInfoModel struct {
	model   domain.Model
	Title   string
	CisInfo domain.CisResult
	MapInfo map[string]interface{}
	Json    string
	Updated time.Time
}

var _ domain.Modeler = (*CisInfoModel)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func NewModel(app domain.Apper) (*CisInfoModel, error) {
	model := &CisInfoModel{
		model:   domain.CisInfo,
		Title:   "Информация по КМ",
		Updated: time.Time{},
		MapInfo: make(map[string]interface{}),
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
func (m *CisInfoModel) SyncToStore(_ domain.Apper) (err error) {
	return err
}

// читаем состояние приложения
func (m *CisInfoModel) ReadState(app domain.Apper) (err error) {
	return nil
}

func (m *CisInfoModel) Copy() (interface{}, error) {
	dst := *m
	// Deep copy the map to avoid shared state
	if m.MapInfo != nil {
		dst.MapInfo = make(map[string]interface{}, len(m.MapInfo))
		for k, v := range m.MapInfo {
			dst.MapInfo[k] = v
		}
	}
	return &dst, nil
}

func (a *CisInfoModel) Model() domain.Model {
	return a.model
}

func (a *CisInfoModel) Save(_ domain.Apper) (err error) {
	return nil
}

// всегда возвращает true означает проверки нет всегда ок
func (m *CisInfoModel) License() bool {
	return true
}
