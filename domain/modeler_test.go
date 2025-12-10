package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelFromText(t *testing.T) {
	model, err := ModelFromString("123")
	assert.NotNil(t, err, "ожидаем ошибку")
	fmt.Println(model)
}

func TestModelFromString_Invalid(t *testing.T) {
	model, err := ModelFromString("123")
	assert.Error(t, err, "ожидаем ошибку")
	assert.Empty(t, model)
}

func TestModel(t *testing.T) {
	model, err := ModelFromString("footer")
	assert.NoError(t, err, "нет ошибки")
	model2 := Model("footer")
	fact := model == model2
	assert.True(t, fact, "ожидаем истину")
	modelErr, err := ModelFromString("123")
	assert.Error(t, err, "ошибка")
	modelErr2 := Model("123")
	fact2 := modelErr == modelErr2
	assert.False(t, fact2, "ожидаем ложь")
}
