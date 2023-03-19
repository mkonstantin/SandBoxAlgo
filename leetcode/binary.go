package leetcode

import (
	"fmt"
	"strconv"
)

type Binary struct{}

func (a *Binary) Start() {
	defer exeTime("Start")()

	fmt.Println(a.findComplement(5))
}

// https://leetcode.com/problems/number-complement/

func (a *Binary) findComplement(num int) int {
	hex := strconv.FormatInt(int64(num), 2)

	runes := []rune(hex)
	for i, rune1 := range runes {
		if rune1 == 49 {
			runes[i] = 48
		} else {
			runes[i] = 49
		}
	}

	marks, err := strconv.ParseInt(string(runes), 2, 0)
	if err != nil {
		return 0
	}

	return int(marks)
}
