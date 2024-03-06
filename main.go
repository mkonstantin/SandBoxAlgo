package main

import (
	"fmt"

	"sandboxProject/stepik"
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
	//learn := concurency.GoroutineOne{}
	//learn.Start()
	//learn := examples.ExamplesOne{}
	//learn.Start()
	//learn := memory.MemoryThree{}
	//learn.Start()
	//learn := arrays.ArrayTwo{}
	//learn.Start()
	//learn := examples.ExamplesTwo{}
	//learn.Start()
	//learn := exp_context.TimeoutContext{}
	//learn.Start()
	learn := stepik.ArraysAndSlices{}
	learn.Start()
}

type StartLearn interface {
	Start()
}
