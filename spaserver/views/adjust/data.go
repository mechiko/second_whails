package adjust

import (
	"fmt"
	"korrectkm/reductor"
)

func (t *page) PageData() (interface{}, error) {
	return reductor.Model[*AdjustModel](t.modelType)
}

// с преобразованием
func (t *page) PageModel() (model *AdjustModel, err error) {
	// model, err := reductor.Instance().Model(t.modelType)
	model, err = reductor.Model[*AdjustModel](t.modelType)
	if err != nil {
		return &AdjustModel{}, fmt.Errorf("%w", err)
	}
	return model, nil
}

// сброс модели редуктора для страницы
func (t *page) ResetData() {
}

func (t *page) ResetValidateData() {
}

func (t *page) ModelUpdate(model *AdjustModel) error {
	err := reductor.SetModel(model, false)
	if err != nil {
		return fmt.Errorf("kmstate page model update %w", err)
	}
	return nil
}
