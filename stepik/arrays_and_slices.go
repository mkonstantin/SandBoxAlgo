package stepik

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type ArraysAndSlices struct {
}

func (s *ArraysAndSlices) Start() {
	lenCapArray()
	sliceLenCap()
	divideString()
	digitsCount()
	phraseText()
	shuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func lenCapArray() {
	var arr [5]int // объявление массива
	//arr := [5]int{1, 2, 3} // создание массива с элементами при объявлении

	// len = cap - так как длина массива известна заранее и равна 5, а items принимают дефолтное значение int = 0
	fmt.Println("Массив", arr)
	fmt.Println("len:", len(arr))
	fmt.Println("cap:", cap(arr))
}

func sliceLenCap() {
	var arr []int // объявление слайса
	//var arr = make([]int, 3) // создание слайса с длиной
	//var arr = make([]int, 3, 5) // создание слайса с длиной (length) и вместимостью (capacity)
	//arr := []int{1, 2, 3} // создание слайса с элементами при объявлении
	arr = append(arr, 1)
	fmt.Println("Срез", arr)
	fmt.Println("len:", len(arr))
	fmt.Println("cap:", cap(arr))
}

func divideString() {
	var text = "wqewefwefwefwef"
	var width = 5

	res := text[0:width]

	fmt.Println(res)
}

func digitsCount() {
	counter := make(map[int]int)

	input := 123143554

	a := strconv.Itoa(input)
	res := strings.Split(a, "")

	for _, re := range res {
		s, _ := strconv.Atoi(re)
		counter[s]++
	}
	fmt.Println(counter)
}

// 1. Разбейте фразу на слова, используя `strings.Fields()`
// 2. Возьмите первую букву каждого слова и приведите
//    ее к верхнему регистру через `unicode.ToUpper()`
// 3. Если слово начинается не с буквы, игнорируйте его
//    проверяйте через `unicode.IsLetter()`
// 4. Составьте слово из получившихся букв и запишите его
//    в переменную `abbr`
// ...

func phraseText() {
	var abbr string
	var phrase = "Hello, 2 world!"

	words := strings.Fields(phrase)
	for _, word := range words {
		for _, w := range word {
			if unicode.IsLetter(w) {
				abbr = abbr + strings.ToUpper(string(w))
				break
			}
		}
	}
	fmt.Println(abbr)
}

func shuffle(nums []int) {
	// перетасуйте nums с помощью rand.Shuffle()
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	fmt.Println(nums)
}

// result представляет результат матча
type result byte

// возможные результаты матча
const (
	win  result = 'W'
	draw result = 'D'
	loss result = 'L'
)

// team представляет команду
type team byte

// match представляет матч
// содержит три поля:
// - first (первая команда)
// - second (вторая команда)
// - result (результат матча)
// например, строке BAW соответствует
// first=B, second=A, result=W
type match struct {
	// ...
	first  team
	second team
	result result
}

// rating представляет турнирный рейтинг команд -
// количество набранных очков по каждой команде
type rating map[team]int

// tournament представляет турнир
type tournament []match

// calcRating считает и возвращает рейтинг турнира
func (trn *tournament) calcRating() rating {
	rating := make(rating)
	for _, elem := range *trn {
		switch elem.result {
		case win:
			rating[elem.first]++
		case draw:
			rating[elem.first]++
			rating[elem.second]++
		case loss:
			rating[elem.second]++
		}
	}
	return rating
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// код, который парсит результаты игр, уже реализован
// код, который печатает рейтинг, уже реализован
// ваша задача - реализовать недостающие структуры и методы выше
func main() {
	src := readString()
	trn := parseTournament(src)
	rt := trn.calcRating()
	rt.print()
}

// readString считывает и возвращает строку из os.Stdin
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}

// parseTournament парсит турнир из исходной строки
func parseTournament(s string) tournament {
	descriptions := strings.Split(s, " ")
	trn := tournament{}
	for _, descr := range descriptions {
		m := parseMatch(descr)
		trn.addMatch(m)
	}
	return trn
}

// parseMatch парсит матч из фрагмента исходной строки
func parseMatch(s string) match {
	return match{
		first:  team(s[0]),
		second: team(s[1]),
		result: result(s[2]),
	}
}

// addMatch добавляет матч к турниру
func (t *tournament) addMatch(m match) {
	*t = append(*t, m)
}

// print печатает результаты турнира
// в формате Aw Bx Cy Dz
func (r *rating) print() {
	var parts []string
	for team, score := range *r {
		part := fmt.Sprintf("%c%d", team, score)
		parts = append(parts, part)
	}
	sort.Strings(parts)
	fmt.Println(strings.Join(parts, " "))
}
