package leetcode

import (
	"fmt"
	"reflect"
	"sort"
)

type BinaryTree struct {
}

func (a *BinaryTree) Start() {

	//fmt.Println(a.SameTree(p, q))
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := -1

	//fmt.Println(a.BinarySearch(nums, target))
	fmt.Println(BinaryRecurse(nums, target))
}

// https://leetcode.com/problems/same-tree/ ============

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var p = &TreeNode{
	Val: 2,
	Left: &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	},
}

var q = &TreeNode{
	Val: 2,
	Left: &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	},
}

// 1 вариант
func (a *BinaryTree) SameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	sdf := ToArray(p)
	arr = nil
	zxc := ToArray(q)
	return reflect.DeepEqual(sdf, zxc)
}

var arr []int

func ToArray(t *TreeNode) []int {
	if t == nil {
		return arr
	}
	ToArray(t.Left)
	arr = append(arr, t.Val)
	ToArray(t.Right)
	return arr
}

// Обход бинарного дерева и вывод узлов

func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}
	t.Left.PrintInorder()
	fmt.Print(t.Val)
	t.Right.PrintInorder()
}

// 2 вариант. Хак
func (a *BinaryTree) SameTree1(p *TreeNode, q *TreeNode) bool {
	return reflect.DeepEqual(p, q)
}

// BinarySearch https://leetcode.com/problems/binary-search/ ==============

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
