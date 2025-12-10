package domain

import (
	"fmt"
	"strings"
)

type Modeler interface {
	Save(Apper) error
	Copy() (interface{}, error) // структура копирует себя и выдает ссылку на копию с массивами и другими данными
	Model() Model               // возвращает тип модели
	License() bool              // если false проверка лицензии не пройдена
}

type Model string

const (
	Application Model = "application"
	TrueClient  Model = "trueclient"
	StatusBar   Model = "statusbar"
	NoPage      Model = "nopage"
	Header      Model = "header"
	Menu        Model = "menu"
	Home        Model = "home"
	Footer      Model = "footer"
	Root        Model = "root"
	KMState     Model = "kmstate"
	InnFias     Model = "innfias"
	Money       Model = "money"
	CisInfo     Model = "cisinfo"
	Adjust      Model = "adjust"
)

func IsValidModel(s string) bool {
	switch Model(s) {
	case Application, TrueClient, StatusBar, NoPage, Header, Footer, Root, Home, KMState, Menu, InnFias, Money, CisInfo, Adjust:
		return true
	default:
		return false
	}
}

// строка приводится в нижний регистр потом сравнивается
func ModelFromString(s string) (Model, error) {
	s = strings.ToLower(s)
	if IsValidModel(s) {
		return Model(s), nil
	}
	return "", fmt.Errorf("%s ошибочная модель domain.Model", s)
}

func (s Model) String() string {
	return string(s)
}
