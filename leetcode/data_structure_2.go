package leetcode

import "fmt"

type DataStructureTwo struct{}

func (a *DataStructureTwo) Start() {
	defer exeTime("Start")()

	fmt.Println(removeDuplicates([]int{1, 2, 3, 4, 4, 5, 6, 7}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates3(nums []int) int {
	for i, num := range nums {
		if i+1 < len(nums) && num == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
	}
	return len(nums)
}

func removeDuplicates(nums []int) int {
	var mapper = make(map[int]int)
	for _, val := range nums {
		mapper[val]++
	}
	nums = nil
	for i, _ := range mapper {
		nums = append(nums, i)
	}
	return len(nums)
}
