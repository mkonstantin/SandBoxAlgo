package leetcode

import (
	"fmt"
	"strings"
)

type StringsOne struct{}

func (a *StringsOne) Start() {
	defer exeTime("Start")()

	fmt.Println(a.lengthOfLastWord("sdf sdkf dsfm m m m    qwwqqq   "))
}

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

func (a *StringsOne) strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// https://leetcode.com/problems/search-insert-position/

func (a *StringsOne) searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

// https://leetcode.com/problems/length-of-last-word/

func (a *StringsOne) lengthOfLastWord(s string) int {
	splitted := strings.Split(strings.TrimRight(s, " "), " ")
	return len(splitted[len(splitted)-1])
}
