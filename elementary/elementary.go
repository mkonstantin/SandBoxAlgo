package elementary

import "fmt"

// инитаем условия

type Human struct {
	Name string
}

func (h *Human) Laugh() {
	fmt.Printf("AHAHAHAHA!!! - %s", h.Name)
}

type Warrior interface {
	Fight()
}

type Archer struct {
	strelaCount int
}

func (w *Archer) Fight() {
	fmt.Println("Fire at will!!!!!")
}

type Man struct {
	Human
	Warrior
	Gun string
}

// запускаем проект
type Elementary struct{}

// Проверяем работу со "Встроенными" типами
// в случаях если этот тип Структура или Интерфейс
func (e *Elementary) Start() {
	man := Man{
		Human: Human{
			Name: "Gans",
		},
		Warrior: &Archer{
			strelaCount: 100,
		},
		Gun: "Colt .50",
	}

	man.Laugh()
	man.Fight()
}
