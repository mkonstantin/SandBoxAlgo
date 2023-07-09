package leetcode

import "fmt"

type Graph1 struct{}

func (a *Graph1) Start() {
	defer exeTime("Start")()

	//[[1,2],[5,1],[1,3],[1,4]]
	//fmt.Println(findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))

	//n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
	//fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))

	//n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
	//fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))

	//n = 10, edges =[[0,7],[0,8],[6,1],[2,0],[0,4],[5,8],[4,7],[1,3],[3,5],[6,5]], source = 7, destination = 5
	//fmt.Println(validPath(10, [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}, 7, 5))

	//[[1,3],[2,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
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

// n = 10, edges =[[0,7],[0,8],[6,1],[2,0],[0,4],[5,8],[4,7],[1,3],[3,5],[6,5]], source = 7, destination = 5
// https://leetcode.com/problems/find-if-path-exists-in-graph/
// Крутое решение по скорости и памяти!!!

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	var queue []int

	queue = append(queue, source)

	for i := 0; i < n; i++ {
		if i >= len(queue) {
			return false
		}
		if queue[i] == destination {
			return true
		} else {
			fri := getFriendsNotContainedInQueue(edges, queue[i], queue)
			queue = append(queue, fri...)
		}
	}
	return false
}

func getFriendsNotContainedInQueue(edges [][]int, dot int, queue []int) []int {
	var s []int
	for _, edge := range edges {
		if edge[0] == dot && !contains(queue, edge[1]) {
			s = append(s, edge[1])
		} else if edge[1] == dot && !contains(queue, edge[0]) {
			s = append(s, edge[0])
		}
	}
	return s
}

func contains(s []int, item int) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

// https://leetcode.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	var ms = make([][]int, 19)
	name := 1

	for _, interval := range intervals {
		ms[interval[0]] = append(ms[interval[0]], name)
		ms[interval[1]] = append(ms[interval[1]], name)
		name++
	}

	var output [][]int
	var out []int

	var queue = make([]int, name)

	for i, m := range ms {
		if m == nil {
			continue
		}

		for _, s := range m {
			queue[s]++
		}

		if out == nil {
			out = append(out, i)
		}
		if checkQueueComplete(queue) {
			out = append(out, i)
			output = append(output, out)
			out = nil
		}
	}
	return output
}

func checkQueueComplete(queue []int) bool {
	for _, item := range queue {
		if item == 1 {
			return false
		}
	}
	return true
}
