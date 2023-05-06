package binarysearch

import "sort"

// BinarySearch https://leetcode.com/problems/binary-search/ ==============
// https://www.code-recipe.com/post/binary-search
// https://www.youtube.com/live/kNkeJ3ZtgJA?feature=share

// Решение 0

func BinarySearchSortLib(nums []int, target int) int {
	i := sort.SearchInts(nums, target)

	if i < len(nums) && nums[i] == target {
		return i
	}

	return -1
}

// Решение 1

func BinaryGo(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for right >= left {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// Recurse func

func BinaryRecurse(arr []int, x int) int {
	// If array is empty return not found(-1)
	if len(arr) == 0 {
		return -1
	}
	return binarySearch(arr, x, 0, len(arr)-1)
}

func binarySearch(arr []int, x, left, right int) int {
	if right >= left {
		mid := (left + right) / 2

		if x == arr[mid] {
			return mid
		} else if x < arr[mid] {
			return binarySearch(arr, x, left, mid-1)
		} else {
			return binarySearch(arr, x, mid+1, right)
		}
	}
	return -1
}
