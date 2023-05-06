package leetcode

import (
	"fmt"
	"math"
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
	fmt.Println(BinarySearchRecurse(nums, target))
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

// https://leetcode.com/problems/binary-search/ ==============
func (a *BinaryTree) BinarySearch(nums []int, target int) int {
	var index = len(nums) / 2

	//var inc = len(nums)
	for index < len(nums) {
		temp := nums[index]
		if target == temp {
			return index
		} else {
			if target > temp {
				index = index + int(math.Round(float64(index)/2))
			} else {
				index = index - int(math.Round(float64(index)/2))
			}
		}
	}
	return -1
}

func BinarySearchRecurse(nums []int, target int) int {
	var index = int(math.Round(float64(len(nums)) / 2))
	temp := nums[index]
	if target == temp {
		return index
	} else {
		if target > temp {
			return BinarySearchRecurse(nums[index:], target)
		} else {
			return BinarySearchRecurse(nums[:index], target)
		}
	}
}

var asd int

func BinarySearchRecurse2(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}
