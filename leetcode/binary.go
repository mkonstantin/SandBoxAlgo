package leetcode

import (
	"fmt"
	"strconv"
)

type Binary struct{}

func (a *Binary) Start() {
	defer exeTime("Start")()

	//fmt.Println(a.findComplement(5))
	//fmt.Println(a.climbStairs(3))
	//fmt.Println(a.maxSubArray([]int{1}))
	fmt.Println(mySqrt(65))
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

// https://leetcode.com/problems/missing-number/description/

func (a *Binary) missingNumber(nums []int) int {
	n := len(nums)

	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum - n*(n+1)/2
}

//https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

func (a *Binary) findDisappearedNumbers(nums []int) []int {
	mapper := make(map[int]int)
	for _, num := range nums {
		mapper[num]++
	}

	var output []int
	for i := 1; i <= len(nums); i++ {
		if mapper[i] == 0 {
			output = append(output, i)
		}
	}

	return output
}

// https://leetcode.com/problems/single-number/

func (a *Binary) singleNumber(nums []int) int {
	var output int
	for _, num := range nums {
		output = output ^ num
	}
	return output
}

// https://leetcode.com/problems/climbing-stairs/

func (a *Binary) climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	n1 := 1
	n2 := 2

	for i := 3; i < n+1; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}

// https://leetcode.com/problems/maximum-subarray/

func (a *Binary) maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum+nums[i] > nums[i] {
			currentSum = currentSum + nums[i]
		} else {
			currentSum = nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// https://leetcode.com/problems/first-bad-version/
func firstBadVersion(n int) int {
	left := 0
	right := n

	for left <= right {
		mid := (left + right) / 2

		if isBadVersion(mid) && !isBadVersion(mid-1) {
			return mid
		} else if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}

func isBadVersion(version int) bool {
	if version >= 10 {
		return true
	}
	return false
}

// https://leetcode.com/problems/valid-perfect-square/

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	left := 0
	right := num / 2

	for left <= right {
		mid := (left + right) / 2

		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// https://leetcode.com/problems/sqrtx/

func mySqrt(x int) int {
	left := 0
	right := x

	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
