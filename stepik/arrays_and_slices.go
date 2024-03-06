package stepik

import (
	"fmt"
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
