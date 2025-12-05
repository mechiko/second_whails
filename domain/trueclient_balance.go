package domain

type Balance []struct {
	Balance        int `json:"balance,omitempty"`
	ContractID     int `json:"contractId,omitempty"`
	OrganisationID int `json:"organisationId"`
	ProductGroupID int `json:"productGroupId"`
}

type GroupNames struct {
	ID    int
	Alias string
	Name  string
}

var ProductGroupIDs = []*GroupNames{
	{ID: 1, Alias: "lp", Name: "Предметы одежды, бельё постельное, столовое, туалетное и кухонное"},
	{ID: 2, Alias: "shoes", Name: "Обувные товары"},
	{ID: 3, Alias: "tobacco", Name: "Табачная продукция"},
	{ID: 4, Alias: "perfumery", Name: "Духи и туалетная вода"},
	{ID: 5, Alias: "tires", Name: "Шины и покрышки пневматические резиновые новые"},
	{ID: 6, Alias: "electronics", Name: "Фотокамеры (кроме кинокамер), фотовспышки и лампывспышки"},
	{ID: 8, Alias: "milk", Name: "Молочная продукция"},
	{ID: 9, Alias: "bicycle", Name: "Велосипеды и велосипедные рамы"},
	{ID: 10, Alias: "wheelchairs", Name: "Медицинские изделия"},
	{ID: 11, Alias: "alcohol", Name: "Алкоголь"},
	{ID: 12, Alias: "otp", Name: "Альтернативная табачная продукция"},
	{ID: 13, Alias: "water", Name: "Упакованная вода"},
	{ID: 14, Alias: "furs", Name: "Товары из натурального меха"},
	{ID: 15, Alias: "beer", Name: "Пиво, напитки, изготавливаемые на основе пива, слабоалкогольные напитки"},
	{ID: 16, Alias: "ncp", Name: "Никотиносодержащая продукция"},
	{ID: 17, Alias: "bio", Name: "Биологически активные добавки к пище"},
	{ID: 19, Alias: "antiseptic", Name: "Антисептики и дезинфицирующие средства"},
	{ID: 20, Alias: "petfood", Name: "Корма для животных"},
	{ID: 21, Alias: "seafood", Name: "Морепродукты"},
	{ID: 22, Alias: "nabeer", Name: "Безалкогольное пиво"},
	{ID: 23, Alias: "softdrinks", Name: "Соковая продукция и безалкогольные напитки"},
	{ID: 25, Alias: "meat", Name: "Мясные изделия"},
	{ID: 26, Alias: "vetpharma", Name: "Ветеринарные препараты"},
	{ID: 27, Alias: "toys", Name: "Игры и игрушки для детей"},
	{ID: 28, Alias: "radio", Name: "Радиоэлектронная продукция"},
	{ID: 31, Alias: "titan", Name: "Титановая металлопродукция"},
	{ID: 32, Alias: "conserve", Name: "Консервированная продукция"},
	{ID: 33, Alias: "vegetableoil", Name: "Растительные масла"},
	{ID: 34, Alias: "opticfiber", Name: "Оптоволокно и оптоволоконная продукция"},
	{ID: 35, Alias: "chemistry", Name: "Косметика, бытовая химия и товары личной гигиены"},
	{ID: 36, Alias: "books", Name: "Печатная продукция"},
	{ID: 37, Alias: "grocery", Name: "Бакалейная продукция"},
	{ID: 38, Alias: "pharmaraw", Name: "Фармацевтическое сырьё, лекарственные средства"},
	{ID: 39, Alias: "construction", Name: "Строительные материалы"},
	{ID: 40, Alias: "fire", Name: "Пиротехника и огнетушащее оборудование"},
	{ID: 41, Alias: "heater", Name: "Отопительные приборы"},
	{ID: 42, Alias: "cableraw", Name: "Кабельно-проводниковая продукция"},
	{ID: 43, Alias: "autofluids", Name: "Моторные масла"},
	{ID: 44, Alias: "polymer", Name: "Полимерные трубы"},
	{ID: 45, Alias: "sweets", Name: "Сладости и кондитерские изделия"},
	{ID: 48, Alias: "carparts", Name: "Автозапчасти и комплектующие транспортных средств"},
	{ID: 50, Alias: "nicotindev", Name: "Радиоэлектронная продукция. Электронные системы доставки никотина"},
	{ID: 51, Alias: "gadgets", Name: "Радиоэлектронная продукция. Ноутбуки и смартфоны"},
}

var ProductGroupByIDs = func() map[int]*GroupNames {
	out := map[int]*GroupNames{}
	for _, gr := range ProductGroupIDs {
		out[gr.ID] = gr
	}
	return out
}()

func ProductGroupByID() map[int]*GroupNames {
	out := map[int]*GroupNames{}
	for _, gr := range ProductGroupIDs {
		out[gr.ID] = gr
	}
	return out
}

func ProductGroupByAlias() map[string]*GroupNames {
	out := map[string]*GroupNames{}
	for _, gr := range ProductGroupIDs {
		out[gr.Alias] = gr
	}
	return nil
}
