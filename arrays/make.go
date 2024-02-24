package arrays

import "fmt"

type ArrayTwo struct{}

func (a *ArrayTwo) Start() {
	Starter()
}

func Starter() {
	arr := []int64{}
	fmt.Printf("\nМассив %#v с %d элементами", arr, len(arr))

	fmt.Println("\nМассив в виде строк:", toModels(arr))
}

func toModels(dtos []int64) []string {
	var arr = make([]string, 0, len(dtos))
	for _, dto := range dtos {
		arr = append(arr, fmt.Sprint("", dto))
	}
	return arr
}
