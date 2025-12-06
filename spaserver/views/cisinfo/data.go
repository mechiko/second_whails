package cisinfo

import (
	"fmt"
	"korrectkm/reductor"
)

func (t *page) PageData() (interface{}, error) {
	return reductor.Model[*CisInfoModel](t.modelType)
}

// с преобразованием
func (t *page) PageModel() (model *CisInfoModel, err error) {
	// model, err := reductor.Instance().Model(t.modelType)
	model, err = reductor.Model[*CisInfoModel](t.modelType)
	if err != nil {
		return &CisInfoModel{}, fmt.Errorf("%w", err)
	}
	return model, nil
}

// сброс модели редуктора для страницы
func (t *page) ResetData() {
}

func (t *page) ResetValidateData() {
}

func (t *page) ModelUpdate(model *CisInfoModel) error {
	err := reductor.SetModel(model, false)
	if err != nil {
		return fmt.Errorf("kmstate page model update %w", err)
	}
	return nil
}
