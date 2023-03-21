package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type DataStructureOne struct{}

func (a *DataStructureOne) Start() {
	defer exeTime("Start")()

	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))

	//merge2([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

	//fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))

	//firstUniqChar("leetcode")

	//fmt.Println(canConstruct("asddd", "dfgasdf"))

	//fmt.Println(isAnagram("ab", "ba"))

	//startListNodeTask()

	//startListNodeTask2()

	//startReverseList()

	//startRemoveList()
}

// func to calculate and print execution time
func exeTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s execution time: %v\n", name, time.Since(start))
	}
}

// https://leetcode.com/problems/contains-duplicate

func containsDuplicate(nums []int) bool {
	defer exeTime("containsDuplicate")()

	var t = make(map[int]int)
	for _, num := range nums {
		if t[num] > 0 {
			return true
		}
		t[num]++
	}
	return false
}

//https://leetcode.com/problems/merge-sorted-array/?envType=study-plan&id=data-structure-i

// первое решение 40ms
func merge1(nums1 []int, m int, nums2 []int, n int) {
	defer exeTime(fmt.Sprint("merge"))()

	nums1 = append(nums1, nums2...)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	nums1 = nums1[n:]

	fmt.Println(nums1)
}

// второе решение быстрее 30ms
func merge2(nums1 []int, m int, nums2 []int, n int) {
	defer exeTime(fmt.Sprint("merge"))()

	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)

	fmt.Println(nums1)
}

// https://leetcode.com/problems/intersection-of-two-arrays-ii/?envType=study-plan&id=data-structure-i

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) <= len(nums2) {
		return containArrays(nums1, nums2)
	} else {
		return containArrays(nums2, nums1)
	}
}

func containArrays(smallArr []int, bigArr []int) []int {
	var output []int
	var outBig = bigArr
	for _, value := range smallArr {
		if sliceContains(value, outBig) {
			if len(output) == len(smallArr) {
				return output
			}
			outBig = removeFromArray(value, outBig)
			output = append(output, value)
		}
	}
	return output
}

func sliceContains(val int, arr []int) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

func removeFromArray(val int, arr []int) []int {
	var out []int
	for i, value := range arr {
		if val == value {
			out = append(arr[:i], arr[i+1:]...)
			return out
		}
	}
	return out
}

//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan&id=data-structure-i

// Рабочий, но долгий неоптимизированный вариант
func maxProfit2(prices []int) int {
	defer exeTime("MaxProfit")()

	var maxRange int
	var arr []int
	for i, price := range prices {
		arr = nil
		arr = append(arr, prices[i+1:]...)
		if len(arr) == 0 {
			return maxRange
		}
		sort.Ints(arr)
		if (arr[len(arr)-1] - price) > maxRange {
			maxRange = arr[len(arr)-1] - price
		}
	}
	return maxRange
}

// Идеальный вариант
func maxProfit(prices []int) int {
	var profit = 0
	var low = 10000
	for _, current := range prices {
		if current < low {
			low = current
			continue
		}
		if profit < (current - low) {
			profit = current - low
		}
	}
	return profit
}

// https://leetcode.com/problems/pascals-triangle/description/?envType=study-plan&id=data-structure-i
// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

//func generate(numRows int) [][]int {
//	var output [][]int
//	for i := 0; i < numRows; i++ {
//
//		dfg := []int{1}
//	}
//}

//https://leetcode.com/problems/first-unique-character-in-a-string/?envType=study-plan&id=data-structure-i

// Мое решение
func firstUniqChar(s string) int {
	for i, val := range s {
		if strings.Count(s, string(val)) == 1 {
			return i
		}
	}
	return -1
}

// Крутое решение с помощью Мапа, посмотрел в ответах
func firstUniqChar2(s string) int {
	nonRepeated := make(map[rune]int)
	for _, char := range s {
		nonRepeated[char]++
	}

	for i, char := range s {
		if nonRepeated[char] == 1 {
			return i
		}
	}

	return -1
}

// https://leetcode.com/problems/ransom-note/?envType=study-plan&id=data-structure-i

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteMap := make(map[rune]int)
	for _, val := range ransomNote {
		ransomNoteMap[val] = ransomNoteMap[val] + 1
	}

	magazineMap := make(map[rune]int)
	for _, val := range magazine {
		magazineMap[val] = magazineMap[val] + 1
	}

	for _, char := range ransomNote {
		if val, ok := magazineMap[char]; ok {
			if ransomNoteMap[char] > val {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/valid-anagram/?envType=study-plan&id=data-structure-i

func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)
	for _, val := range s {
		sMap[val] = sMap[val] + 1
	}

	tMap := make(map[rune]int)
	for _, val := range t {
		tMap[val] = tMap[val] + 1
	}

	if len(sMap) != len(tMap) {
		return false
	}

	for _, char := range s {
		val, ok := tMap[char]
		if !ok {
			return false
		}
		if sMap[char] != val {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/linked-list-cycle/description/?envType=study-plan&id=data-structure-i

var node1 = ListNode{
	Val:  3,
	Next: nil,
}

var node2 = ListNode{
	Val:  2,
	Next: nil,
}

var node3 = ListNode{
	Val:  0,
	Next: nil,
}

var node4 = ListNode{
	Val:  -4,
	Next: nil,
}

func startListNodeTask() {
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node2

	fmt.Println(hasCycle(&node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}

	var pos = 0
	var current = head
	var next = current.Next
	var arr = make(map[*ListNode]int)

	arr[current] = pos

	for next != nil {
		next = current.Next
		if _, ok := arr[next]; ok {
			return true
		}
		pos++
		arr[next] = pos
		current = next
	}
	return false
}

//https://leetcode.com/problems/merge-two-sorted-lists/?envType=study-plan&id=data-structure-i

var node10 = ListNode{
	Val:  -10,
	Next: nil,
}

var node20 = ListNode{
	Val:  -10,
	Next: nil,
}

var node30 = ListNode{
	Val:  -9,
	Next: nil,
}

var node40 = ListNode{
	Val:  -4,
	Next: nil,
}

var node50 = ListNode{
	Val:  1,
	Next: nil,
}

var node60 = ListNode{
	Val:  6,
	Next: nil,
}

var node70 = ListNode{
	Val:  6,
	Next: nil,
}

var node80 = ListNode{
	Val:  -7,
	Next: nil,
}

func startListNodeTask2() {
	node10.Next = &node20
	node20.Next = &node30
	node30.Next = &node40
	node40.Next = &node50
	node50.Next = &node60
	node60.Next = &node70
	node70.Next = nil

	node80.Next = nil

	fmt.Println(mergeTwoLists(&node10, &node80))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var arr = make(map[int]*ListNode)

	if list1 != nil {
		arr = addToArr(arr, list1)
	}
	if list2 != nil {
		arr = addToArr(arr, list2)
	}
	if len(arr) == 0 {
		return nil
	}
	var intList []*ListNode
	for _, val := range arr {
		val.Next = nil
		intList = append(intList, val)
	}
	sort.Slice(intList, func(p, q int) bool {
		return intList[p].Val < intList[q].Val
	})

	for i, node := range intList {
		if i+1 == len(intList) {
			break
		}
		node.Next = intList[i+1]
	}

	return intList[0]
}

func addToArr(arr map[int]*ListNode, current *ListNode) map[int]*ListNode {
	var pos = len(arr)
	for current != nil {
		arr[pos] = current
		current = current.Next
		pos++
	}
	return arr
}

//https://leetcode.com/problems/reverse-linked-list/?envType=study-plan&id=data-structure-i

var node11 = ListNode{
	Val:  1,
	Next: nil,
}

var node21 = ListNode{
	Val:  2,
	Next: nil,
}

var node31 = ListNode{
	Val:  3,
	Next: nil,
}

var node41 = ListNode{
	Val:  4,
	Next: nil,
}

var node51 = ListNode{
	Val:  5,
	Next: nil,
}

func startReverseList() {
	node11.Next = &node21
	node21.Next = &node31
	node31.Next = &node41
	node41.Next = &node51
	node51.Next = nil

	fmt.Println(reverseList(&node11))
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	if head.Next == nil {
//		return head
//	}
//
//	mappedList := convertToMap(head)
//
//	end := mappedList[len(mappedList)-1]
//
//	for i := len(mappedList) - 1; i > 0; i-- {
//		mappedList[i].Next = mappedList[i-1]
//	}
//
//	return end
//}
//
//func convertToMap(current *ListNode) map[int]*ListNode {
//	var arr = make(map[int]*ListNode)
//	var pos = 0
//
//	for current != nil {
//		arr[pos] = current
//		current = current.Next
//		pos++
//	}
//
//	for _, node := range arr {
//		node.Next = nil
//	}
//	return arr
//}

func helper(current *ListNode, prev *ListNode) *ListNode {
	if current == nil {
		return prev
	}
	next := current.Next
	current.Next = prev
	return helper(next, current)
}

func reverseList(head *ListNode) *ListNode {
	return helper(head, nil)
}

//https://leetcode.com/problems/remove-linked-list-elements/description/

var node101 = ListNode{
	Val:  1,
	Next: nil,
}

var node201 = ListNode{
	Val:  2,
	Next: nil,
}

var node301 = ListNode{
	Val:  3,
	Next: nil,
}

var node401 = ListNode{
	Val:  3,
	Next: nil,
}

var node501 = ListNode{
	Val:  5,
	Next: nil,
}

func startRemoveList() {
	node101.Next = &node201
	node201.Next = &node301
	node301.Next = &node401
	node401.Next = nil
	//node501.Next = nil

	dfg := deleteDuplicates(&node101)
	fmt.Println(dfg)
}

// Первое решение скорость: О(1) память О(n)

func removeElements1(head *ListNode, val int) *ListNode {
	var sentinel = &ListNode{Next: head}

	for curr, currNext := sentinel, head; currNext != nil; {
		if currNext.Val == val {
			curr.Next, currNext = currNext.Next, currNext.Next
		} else {
			curr, currNext = curr.Next, currNext.Next
		}
	}
	return sentinel.Next
}

// Первое решение скорость: О(n) память О(n)

func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElements2(head.Next, val)

	if head.Val == val {
		return head.Next
	}

	return head
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/?envType=study-plan&id=data-structure-i

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Val == head.Next.Val {
		head = deleteDuplicates(head.Next)
		return head
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}
