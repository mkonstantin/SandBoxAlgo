package leetcode

import (
	"fmt"
	"sort"
)

type Binary2 struct{}

func (a *Binary2) Start() {
	defer exeTime("Start")()

	//fmt.Println(a.findComplement(5))
	//fmt.Println(a.climbStairs(3))
	//fmt.Println(a.maxSubArray([]int{1}))
	//fmt.Println(intersection([]int{4, 9, 5, 4, 1}, []int{9, 4, 9, 8, 4, 6, 7}))
	//fmt.Println(guessNumber(20))
	fmt.Println(peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))
}

//https://leetcode.com/problems/find-the-duplicate-number/
func findDuplicate(nums []int) int {
	sort.Ints(nums)

	prev := 0
	for _, num := range nums {
		if num == prev {
			return num
		}
		prev = num
		//if search(nums, num) {
		//	prev = num
		//}
	}
	return 0
}

// https://leetcode.com/problems/peak-index-in-a-mountain-array/

func peakIndexInMountainArray(arr []int) int {
	var mapper = make(map[int]int)

	for index, item := range arr {
		i := search(arr, item)
		if i >= 0 {
			mapper[i] = index
		}
	}

	arr = nil
	for _, i2 := range mapper {
		arr = append(arr, i2)
	}
	return arr[len(arr)-1]
}

func search(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
