package examples

import (
	"fmt"
	"unicode/utf8"

	"sandboxProject/pkg"
)

type sdfsdf interface{}

type ExamplesOne struct{}

func (a *ExamplesOne) Start() {
	defer pkg.ExecuteTime("Start")()

	//GOMAXPROCSExp()
	//WaitGroupExp()

	StringsArray()
}

func StringsArray() {
	s := "Boat 🎂."

	fmt.Println("lenght:", len(s))

	fmt.Println("rune count:", utf8.RuneCount([]byte(s)))

	fmt.Println("rune len of A:", utf8.RuneLen([]rune("A")[0]))

	fmt.Println("rune len of 🎂:", utf8.RuneLen([]rune("🎂")[0]))

	// string - массив байтов (8-битное представление или utf-8) т е когда мы вызываем len(s) то получаем кол-во байт (8 бит)
	// В целом строку по-СИМВОЛЬНО можно представить как массив rune. Rune - это alias int32. т е может занимать до 4 байт (32 бита)
	// Если у нас в строке встречается большой символ, то он модет быть представлен как 1-2-3-4 байта. Выяснить жто можно с помощью utf8.RuneLen
	// Чтобы узнать сколько именно символов в строке, то юзать надо функцию utf8.RuneCount
	for i, c := range s {
		fmt.Printf("position %d of %s \n", i, string(c))
	}
}
