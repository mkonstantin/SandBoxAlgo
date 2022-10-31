package main

import (
	"fmt"
	"sandboxProject/elementary"
)

func main() {
	fmt.Println("Start program")

	learn := elementary.Elementary{}
	learn.Start()
}

type StartLearn interface {
	Start()
}
