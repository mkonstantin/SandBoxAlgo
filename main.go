package main

import (
	"fmt"
	"sandboxProject/leetcode"
)

func main() {
	fmt.Println("Start program")

	//learn := elementary.Elementary{}
	//learn := panic_lesson.PanicLesson{}
	//learn := types.TypesExamples{}
	//learn := leetcode.LessonsOne{}
	//learn.Start()
	//learn := leetcode.DataStructureOne{}
	//learn.Start()
	learn := leetcode.DataStructureTwo{}
	learn.Start()
}

type StartLearn interface {
	Start()
}
