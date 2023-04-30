package leetcode

import "fmt"

type AlgorithmOne struct {
}

func (a *AlgorithmOne) Start() {
	//array := a.Constructor([]int{-2, 0, 3, -5, 2, -1})
	//fmt.Println(array.sumRange(0, 2))
	//fmt.Println(array.sumRange(2, 5))
	//fmt.Println(array.sumRange(0, 5))

	fmt.Println(a.countBits(5))
}

// https://leetcode.com/problems/range-sum-query-immutable/

type NumArray struct {
	array []int
	sums  []int
}

func (a *AlgorithmOne) Constructor(nums []int) NumArray {
	var sums []int
	var sum int
	for _, num := range nums {
		sum = sum + num
		sums = append(sums, sum)
	}
	numArray := NumArray{
		array: nums,
		sums:  sums,
	}
	return numArray
}

func (n *NumArray) sumRange(left int, right int) int {
	if left == 0 {
		return n.sums[right]
	}
	return n.sums[right] - n.sums[left-1]
}

// https://leetcode.com/problems/counting-bits/

func (a *AlgorithmOne) countBits(n int) []int {
	arr := make([]int, n+1, n+1)

	for i := 0; i <= n; i++ {
		arr[i] = arr[i>>1] + i%2
	}

	return arr
}
