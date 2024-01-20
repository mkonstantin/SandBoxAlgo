package main

import (
	"fmt"

	"sandboxProject/concurency"
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
	//learn := leetcode.Binary{}
	//learn.Start()
	//learn := leetcode.StringsOne{}
	//learn.Start()
	//learn := leetcode.ArrayOne{}
	//learn.Start()
	//learn := leetcode.NumberOne{}
	//learn.Start()
	//learn := leetcode.AlgorithmOne{}
	//learn.Start()
	//learn := leetcode.Binary2{}
	//learn.Start()
	//learn := leetcode.Graph1{}
	//learn.Start()
	//learn := leetcode.BinaryTree{}
	//learn.Start()
	//learn := leetcode.StringsOne{}
	//learn.Start()

	learn := concurency.GoroutineOne{}
	learn.Start()
}

type StartLearn interface {
	Start()
}
