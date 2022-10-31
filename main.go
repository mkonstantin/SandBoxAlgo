package main

import (
	"fmt"
	"sandboxProject/panic_lesson"
)

func main() {
	fmt.Println("Start program")

	//learn := elementary.Elementary{}
	learn := panic_lesson.PanicLesson{}
	learn.Start()
}

type StartLearn interface {
	Start()
}
