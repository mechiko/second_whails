package gtin

import (
	"fmt"
	"korrectkm/reductor"
)

func (t *page) PageData() (interface{}, error) {
	return reductor.Model[*TargetModel](t.modelType, t)
}

// с преобразованием
func (t *page) PageModel() (model *TargetModel, err error) {
	// model, err := reductor.Instance().Model(t.modelType)
	model, err = reductor.Model[*TargetModel](t.modelType, t)
	if err != nil {
		return &TargetModel{}, fmt.Errorf("%w", err)
	}
	return model, nil
}

// сброс модели редуктора для страницы
func (t *page) ResetData() {
}

func (t *page) ResetValidateData() {
}

func (t *page) ModelUpdate(model *TargetModel) error {
	err := reductor.SetModel(model, false)
	if err != nil {
		return fmt.Errorf("target page model update %w", err)
	}
	return nil
}
