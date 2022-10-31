package panic_lesson

import "fmt"

type PanicLesson struct{}

func (p *PanicLesson) Start() {
	//lessonOne()
	lessonTwo()
}

// Primer 1
// функции defer всегда вызываются после паники
// результат будет:
// "go  now step 1" -> "defer function run"
func lessonOne() {
	defer func() {
		fmt.Println("defer function run")
	}()

	fmt.Println("go now step 1-1")
	panic("we got a panic")
}

// Урок 2 Паникуем и восстанавливаемся
// Вызываем панику, в Дефер функции вызываем recover()
// И дальше продолжаем работу пропуская функцию с паникой
func lessonTwo() {
	fmt.Println("Шаг 1: вызываем функцию с паникой")
	PanicFunc()
	fmt.Println("Шаг 4: восстановились после паники как ни в чем не бывало")
}

func PanicFunc() {
	defer func() {
		fmt.Println("defer function run")
		fmt.Println("Шаг 3: Начинаем восстанавливаться")
		recover()
	}()

	fmt.Println("Шаг 2: Паникуем")
	panic("we got a panic")
}
