package leetcode

import "fmt"

type NumberOne struct{}

func (a *NumberOne) Start() {
	fmt.Println(a.isPowerOfThree(1162261467))
}

// https://leetcode.com/problems/power-of-three/

func (a *NumberOne) isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	flo := float64(n)
	for flo > 1 {
		flo = flo / 3
	}
	return flo == 1
}
