package leetcode

import "fmt"

type DataStructureTwo struct{}

func (a *DataStructureTwo) Start() {
	defer exeTime("Start")()

	//fmt.Println(a.removeDuplicates([]int{1, 2, 3, 4, 4, 5, 6, 7}))
	//fmt.Println(a.removeDuplicates([]int{1, 1}))
	//fmt.Println(a.removeDuplicates([]int{0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(a.removeElement([]int{0, 1, 1, 1, 2, 2, 3, 3, 4}, 1))
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func (a *DataStructureTwo) removeDuplicates3(nums []int) int {
	for i, num := range nums {
		if i+1 < len(nums) && num == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
	}
	return len(nums)
}

func (a *DataStructureTwo) removeDuplicates(nums []int) int {
	defer exeTime("removeDuplicates")()

	var mapper = make(map[int]int)
	for _, val := range nums {
		mapper[val]++
	}
	for i, num := range nums {
		if mapper[num] > 1 {
			nums = append(nums[:i], nums[i+mapper[num]-1:]...)
			delete(mapper, num)
			continue
		}
	}

	return len(nums)
}

func (a *DataStructureTwo) removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
