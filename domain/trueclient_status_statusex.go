package domain

import (
	"fmt"
	"strings"
)

type Status struct {
	ID     int
	Alias  string
	Name   string
	Actual string
}

var StatusList = []Status{
	{ID: 0, Alias: "EMITTED", Name: "Эмитирован"},
	{ID: 1, Alias: "APPLIED", Name: "Нанесён"},
	{ID: 2, Alias: "INTRODUCED", Name: "В обороте"},
	{ID: 3, Alias: "WRITTEN_OFF", Name: "Списан"},
	{ID: 4, Alias: "RETIRED", Name: "Выбыл"},
	{ID: 4, Alias: "WITHDRAWN", Name: "Выбыл"},
	{ID: 7, Alias: "DISAGGREGATION", Name: "Расформирован"},
	{ID: 7, Alias: "DISAGGREGATED", Name: "Расформирован"},
	{ID: 12, Alias: "APPLIED_NOT_PAID", Name: "Не оплачен"},
}

var StatusNameByAlias = func() map[string]string {
	out := map[string]string{}
	for _, gr := range StatusList {
		out[strings.ToLower(gr.Alias)] = gr.Name
	}
	return out
}()

var StatusExList = []Status{
	{Alias: "IN_GRAY_ZONE", Name: "В Серой зоне", Actual: "APPLIED"},
	{Alias: "EMPTY", Name: "отсутствует", Actual: ""},
	{Alias: "RESERVED_NOT_USED", Name: "Зарезервировано. Не использовать", Actual: "INTRODUCED"},
	{Alias: "INDIVIDUAL", Name: "КиЗ индивидуализирован", Actual: "EMITTED, APPLIED"},
	{Alias: "NON_INDIVIDUAL", Name: "Не индивидуализирован", Actual: "EMITTED, APPLIED"},
	{Alias: "WAIT_SHIPMENT", Name: "Ожидает приёмку товара", Actual: "INTRODUCED"},
	{Alias: "EXPORTED", Name: "Используется для документов экспорта", Actual: ""},
	{Alias: "LOAN_RETIRED", Name: "Выведен из оборота по договору рассрочки", Actual: "RETIRED"},
	{Alias: "LOST_INVENTORY", Name: "Не найдены по итогу инвентаризации", Actual: "INTRODUCED"},
	{Alias: "REMARK_RETIRED", Name: "Перемаркирован", Actual: "WRITTEN_OFF"},
	{Alias: "WAIT_TRANSFER_TO_OWNER", Name: "Ожидает передачу собственнику", Actual: "INTRODUCED"},
	{Alias: "WAIT_REMARK", Name: "Ожидает перемаркировку", Actual: "WRITTEN_OFF"},
	{Alias: "RETIRED_CANCELLATION", Name: "Списан / Аннулирован", Actual: "WRITTEN_OFF"},
	{Alias: "FTS_RESPOND_NOT_OK", Name: "Отрицательное решение ФТС", Actual: "APPLIED"},
	{Alias: "FTS_RESPOND_WAITING", Name: "Ожидает подтверждение ФТС", Actual: "APPLIED"},
	{Alias: "FTS_CONTROL", Name: "На контроле ФТС", Actual: "APPLIED"},
	{Alias: "EAS_RESPOND_NOT_OK", Name: "Отрицательное решение ЕАЭС", Actual: "RETIRED"},
	{Alias: "EAS_RESPOND_WAITING", Name: "Ожидает подтверждение ЕАЭС", Actual: "RETIRED"},
	{Alias: "CONNECT_TAP", Name: "Подключён к оборудованию для розлива", Actual: "INTRODUCED, RETIRED"},
	{Alias: "PRIM_RESPONSE_WAITING", Name: "Обрабатывается", Actual: "INTRODUCED, APPLIED, RETIRED"},
	{Alias: "MOVING_BY_UD", Name: "Отгружен", Actual: "INTRODUCED"},
}

var StatusExNameByAlias = func() map[string]string {
	out := map[string]string{}
	for _, gr := range StatusExList {
		if gr.Actual != "" {
			out[strings.ToLower(gr.Alias)] = fmt.Sprintf("%s [%s]", gr.Name, gr.Actual)
		} else {
			out[strings.ToLower(gr.Alias)] = gr.Name
		}
	}
	return out
}()
