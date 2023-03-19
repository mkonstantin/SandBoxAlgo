package leetcode

import (
	"fmt"
	"strings"
)

type LessonsOne struct{}

func (a *LessonsOne) Start() {
	//var array = []int{3, 3}
	//twoSum := twoSum(array, 6)
	//fmt.Println(twoSum)

	//boole := isPalindrome(100)
	//fmt.Println(boole)

	//fmt.Println(romanToInteger("IV"))

	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//
	//fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}

func twoSum(nums []int, target int) []int {
	for index, val := range nums {
		sumIndex := getSumIndex(nums, val, index, target)
		if sumIndex != 0 {
			return []int{index, sumIndex}
		}
	}
	return []int{0, 0}
}

func getSumIndex(nums []int, compareVal int, compareIndex int, target int) int {
	for index, val := range nums {
		if compareIndex == index {
			continue
		}
		sum := compareVal + val
		if sum == target {
			return index
		}
	}
	return 0
}

// https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return true
}

// https://leetcode.com/problems/roman-to-integer/

var numbersMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

type RomanInt struct {
	roman string
	value int
}

var rules = []RomanInt{
	{roman: "IV", value: 4},
	{roman: "IX", value: 9},
	{roman: "XL", value: 40},
	{roman: "XC", value: 90},
	{roman: "CD", value: 400},
	{roman: "CM", value: 900},
}

func romanToInteger(romanStr string) int {

	var findedComplex = make(map[int]RomanInt)
	for _, val := range rules {
		if i := strings.Index(romanStr, val.roman); i >= 0 {
			findedComplex[i] = val
		}
	}

	var outputArr []RomanInt
	var flag bool
	romanArr := strings.Split(romanStr, "")
	for index, val := range romanArr {
		if flag {
			flag = false
			continue
		}
		if compl, ok := findedComplex[index]; ok {
			flag = true
			outputArr = append(outputArr, compl)
		} else {
			number := numbersMap[val]
			ri := RomanInt{roman: val, value: number}
			outputArr = append(outputArr, ri)
		}
	}

	var sum int
	for _, val := range outputArr {
		sum = sum + val.value
	}

	return sum
}

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	startWord := getSmallestString(strs)

	for i := len(startWord); i > 0; i-- {
		prefix := startWord[0:i]
		if isCommonPrefix(prefix, strs) {
			return prefix
		}
	}

	return ""
}

func isCommonPrefix(prefix string, strs []string) bool {
	for _, str := range strs {
		if !strings.HasPrefix(str, prefix) {
			return false
		}
	}
	return true
}

func getSmallestString(strs []string) string {
	smallest := strs[0]
	for _, val := range strs {
		if len(smallest) > len(val) {
			smallest = val
		}
	}
	return smallest
}
