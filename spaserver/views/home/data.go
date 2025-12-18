package home

import (
	"korrectkm/reductor"
)

func (t *page) PageData() (interface{}, error) {
	return reductor.Model[*HomeModel](t.modelType, t)
}

// с преобразованием
func (t *page) PageModel() HomeModel {
	model, _ := reductor.Model[*HomeModel](t.modelType, t)
	return *model
}

// сброс модели редуктора для страницы
func (t *page) ResetData() {
}

func (t *page) ResetValidateData() {
}
