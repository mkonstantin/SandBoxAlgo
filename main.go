package main

import (
	"fmt"
	"sandboxProject/arrays"
)

func main() {
	fmt.Println("Start program")

	//learn := elementary.Elementary{}
	//learn := panic_lesson.PanicLesson{}
	learn := arrays.ArraysExamples{}
	learn.Start()
}

type StartLearn interface {
	Start()
}
