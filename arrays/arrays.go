package arrays

import "fmt"

type ArraysExamples struct{}

func (a *ArraysExamples) Start() {
	lessonThree()
	lessonFour()
	lessonFive()
}

func lessonOne() {
	var variab1 int
	ans := fmt.Sprintf("\nЧисло: %d", variab1)
	fmt.Println(ans)
}

func lessonTwo() {
	var array [5]string
	array[0] = "qwe"
	fmt.Printf("\nМассив %#v с %d элементами", array, len(array))
}

func lessonThree() {
	array := [5]string{"qwe1", "qwe2", "qwe3"}
	array[0] = "qwe100"
	fmt.Printf("\nМассив %#v с %d элементами", array, len(array))
}

func lessonFour() {
	var array = [3]int{1, 2, 3}
	fmt.Printf("\nМассив %#v", array)
}

func lessonFive() {
	//var mymap = map[string]string{} // - инициализировал 1 способ
	//mymap := make(map[string]string) // - инициализировал 2 способ
	//var mymap map[string]string     // - объявил переменную и создал с помощью make 3 способ
	//mymap = make(map[string]string) // |
	mymap := map[string]string{"a1": "A1", "a2": "A2", "a3": "A3"}
	mymap["asd"] = "fsg"

	for key, value := range mymap {
		fmt.Printf("\nkey: %s value: %s", key, value)
	}
}

// Таким образом можно проверять есть ли значение в переменной карты или его нет
// при обращении к ключу карты ok покажет, есть такой ключ или его нет
func lesson6() {
	mymap := make(map[string]string)
	mymap["a1"] = "A1"
	mymap["a2"] = "A2"

	value, ok := mymap["a1"] // тут есть
	fmt.Printf("\nvalue: %s isExist: %v", value, ok)

	value, ok = mymap["DDD"] // такого ключа нет
	fmt.Printf("\nvalue: %s isExist: %v", value, ok)
}
