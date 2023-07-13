package leetcode

import (
	"fmt"
)

type ArrayOne struct{}

func (a *ArrayOne) Start() {
	//fmt.Println(a.plusOne([]int{9, 9}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
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

// https://leetcode.com/problems/majority-element
func majorityElement(nums []int) int {
	mp := make(map[int]float32)
	df := float32(len(nums)) / 2

	for _, num := range nums {
		mp[num]++
		if mp[num] > df {
			return num
		}
	}
	return 0
}
