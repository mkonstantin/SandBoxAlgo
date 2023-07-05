package leetcode

import "fmt"

type Graph1 struct{}

func (a *Graph1) Start() {
	defer exeTime("Start")()

	//[[1,2],[5,1],[1,3],[1,4]]
	fmt.Println(findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))
}

// https://leetcode.com/problems/find-the-town-judge/
// СУПЕР! Вычисляет сколько человек доверяют человеку + этот человек никому не верит!
func findJudge(n int, trust [][]int) int {
	inDegree := make([]int, n+1)
	outDegree := make([]int, n+1)
	for _, t := range trust {
		inDegree[t[1]]++
		outDegree[t[0]]++
	}
	for i := 1; i <= n; i++ {
		if inDegree[i] == n-1 && outDegree[i] == 0 {
			return i
		}
	}
	return -1
}

//func findJudge(n int, trust [][]int) int {
//
//	if len(trust) == 0 && n > 1 {
//		return -1
//	}
//
//	queue := make(map[int][]int)
//
//	for i := 1; i <= n; i++ {
//		for _, ints := range trust {
//			if i == ints[0] {
//				queue[i] = append(queue[i], ints[1])
//			}
//		}
//	}
//
//	temp := 0
//	for i := 1; i <= n; i++ {
//		_, ok := queue[i]
//		if !ok {
//			temp = i
//		}
//	}
//	if temp == 0 {
//		return -1
//	}
//
//	for _, ints := range queue {
//		exist := false
//		for _, dig := range ints {
//			if temp == dig {
//				exist = true
//			}
//		}
//		if !exist {
//			return -1
//		}
//	}
//	return temp
//}

// https://leetcode.com/problems/find-center-of-star-graph/
// Массив действует быстрее мапы, за счет инициализации мы знали сколько будет узлов, это позволило сохранять сразу по нужному индексу, что дало прирост в скорости
func findCenter(edges [][]int) int {
	sorted := make([]int, len(edges)+2)

	for _, edge := range edges {
		sorted[edge[0]]++
		sorted[edge[1]]++
	}

	answer := 0
	max := 0
	for i, item := range sorted {
		if item > max {
			max = item
			answer = i
		}
	}
	return answer
}
