package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type StringsOne struct{}

func (a *StringsOne) Start() {
	defer exeTime("Start")()

	//fmt.Println(a.lengthOfLastWord("sdf sdkf dsfm m m m    qwwqqq   "))
	//fmt.Println(a.repeatedSubstringPattern("a"))
	//fmt.Println(isValid("[]{}()"))
	SortedExample()
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

// https://leetcode.com/problems/repeated-substring-pattern/

func (a *StringsOne) repeatedSubstringPattern(trimmed string) bool {
	for i := 1; i < len(trimmed); i++ {
		if strings.Count(trimmed, trimmed[0:i])*len(trimmed[0:i]) == len(trimmed) {
			return true
		}
	}
	return false
}

// https://leetcode.com/problems/reverse-string/

func reverseString(s []byte) {
	z := len(s) - 1

	for i := 0; i < len(s); i++ {
		if i == z {
			break
		}
		g := s[i]
		s[i] = s[z]
		s[z] = g
		z--
	}
	fmt.Println(string(s))
}

// https://leetcode.com/problems/valid-parentheses/
// "()[]{}"
// "()"
// "()[]{}"
// "(]"
// "[()]"
// "([)]"

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	array := strings.Split(s, "")

	if len(array)%2 != 0 {
		return false
	}

	if array[0] == "]" || array[0] == "}" || array[0] == ")" {
		return false
	}

	for i, item := range array {
		if i == len(array)-1 || i == 0 {
			continue
		}
		switch item {
		case "{":
			if array[i+1] == "}" {
				temp := array[:i]
				temp2 := array[i+2:]
				array = append(temp, temp2...)
			}
		case "[":
			if array[i+1] == "]" {
				temp := array[:i]
				temp2 := array[i+2:]
				array = append(temp, temp2...)
			}
		case "(":
			if array[i+1] == ")" {
				temp := array[:i]
				temp2 := array[i+2:]
				array = append(temp, temp2...)
			}
		}
	}

	return true
}

func isValid2(s string) bool {
	if len(s) <= 1 {
		return false
	}

	array := strings.Split(s, "")

	if len(array)%2 != 0 {
		return false
	}

	if array[0] == "]" || array[0] == "}" || array[0] == ")" {
		return false
	}

	ms := make(map[string]int)

	for _, item := range array {
		ms[item]++
	}

	if ms["{"] != ms["}"] {
		return false
	}
	if ms["["] != ms["]"] {
		return false
	}
	if ms["("] != ms[")"] {
		return false
	}

	if checker(array) {
		return true
	}

	j := len(array) - 1
	for i := 0; i < (j+1)/2; i++ {
		switch array[i] {
		case "{":
			if array[j-i] != "}" {
				return false
			}
		case "[":
			if array[j-i] != "]" {
				return false
			}
		case "(":
			if array[j-i] != ")" {
				return false
			}
		case "}":
			return false
		case "]":
			return false
		case ")":
			return false
		}
	}

	return true
}

func checker(array []string) bool {
	for i, item := range array {
		if i == len(array)-1 || i == 0 {
			continue
		}
		switch item {
		case "{":
			if array[i+1] != "}" {
				return false
			}
		case "[":
			if array[i+1] != "]" {
				return false
			}
		case "(":
			if array[i+1] != ")" {
				return false
			}
		case "}":
			if array[i-1] != "{" {
				return false
			}
		case "]":
			if array[i-1] != "[" {
				return false
			}
		case ")":
			if array[i-1] != "(" {
				return false
			}
		}

	}
	return true
}

type Item struct {
	Weight    int
	CreatedAt time.Time
}

type ByWeightAndCreatedAt []Item

func (a ByWeightAndCreatedAt) Len() int      { return len(a) }
func (a ByWeightAndCreatedAt) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Функция Less определяет порядок сортировки.
// Если значения "weight" у двух элементов различаются, сортировка производится по "weight".
// Если значения "weight" одинаковы, тогда сортировка производится по "created_at".
func (a ByWeightAndCreatedAt) Less(i, j int) bool {
	if a[i].Weight != a[j].Weight {
		return a[i].Weight > a[j].Weight
	}
	return a[i].CreatedAt.After(a[j].CreatedAt)
}

func SortedExample() {
	// Пример среза данных
	items := []Item{
		{Weight: 3, CreatedAt: time.Date(2023, time.July, 25, 10, 30, 0, 0, time.UTC)},
		{Weight: 1, CreatedAt: time.Date(2023, time.July, 25, 9, 0, 0, 0, time.UTC)},
		{Weight: 2, CreatedAt: time.Date(2023, time.July, 25, 11, 15, 0, 0, time.UTC)},
		{Weight: 2, CreatedAt: time.Date(2023, time.July, 25, 10, 0, 0, 0, time.UTC)},
		{Weight: 1, CreatedAt: time.Date(2023, time.July, 25, 8, 0, 0, 0, time.UTC)},
		{Weight: 2, CreatedAt: time.Date(2023, time.July, 25, 12, 0, 0, 0, time.UTC)},
		{Weight: 3, CreatedAt: time.Date(2023, time.July, 24, 10, 0, 0, 0, time.UTC)},
		{Weight: 3, CreatedAt: time.Date(2023, time.July, 26, 18, 0, 0, 0, time.UTC)},
	}

	// Сортировка среза по двум полям: сначала по "weight", затем по "created_at"
	//sort.Sort(ByWeightAndCreatedAt(items))

	sort.Slice(items, func(i int, j int) bool {
		if items[i].Weight != items[j].Weight {
			return items[i].Weight > items[j].Weight
		}
		return items[i].CreatedAt.After(items[j].CreatedAt)
	})

	// Вывод отсортированных элементов
	for _, item := range items {
		fmt.Printf("Weight: %d, CreatedAt: %s\n", item.Weight, item.CreatedAt)
	}
}
