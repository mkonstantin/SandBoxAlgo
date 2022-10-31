package panic_lesson

import "fmt"

type PanicLesson struct{}

func (p *PanicLesson) Start() {
	lessonOne()
}

// Primer
// функции defer всегда вызываются после паники
// результат будет:
// "go  now step 1" -> "defer function run"
func lessonOne() {
	defer func() {
		fmt.Println("defer function run")
	}()

	fmt.Println("go  now step 1")
	panic("we got a panic")
}
