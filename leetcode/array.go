package leetcode

import (
	"fmt"
)

type ArrayOne struct{}

func (a *ArrayOne) Start() {
	fmt.Println(a.plusOne([]int{9, 9}))
}

func (a *ArrayOne) plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			continue
		}
		digits[i] = digits[i] + 1
		return digits
	}
	return append([]int{1}, digits...)
}
