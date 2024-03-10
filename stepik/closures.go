package stepik

import "fmt"

type Closures struct{}

func (s *Closures) Start() {
	filtersPredicate()
}

func filter(predicate func(int) bool, iterable []int) []int {
	var res []int
	for _, item := range iterable {
		if predicate(item) {
			res = append(res, item)
		}
	}
	return res
}

func filtersPredicate() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := filter(func(i int) bool {
		return i%2 == 0
	}, arr)
	fmt.Println(res)
}
