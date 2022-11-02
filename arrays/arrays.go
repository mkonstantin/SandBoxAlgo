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
	var mymap = map[string]string{} // - инициализировал
	//mymap := make(map[string]string) // - инициализировал
	//var mymap map[string]string // - просто объявил переменную

	mymap["asd"] = "fsg"
	for key, value := range mymap {
		fmt.Printf("\nkey: %s value: %s", key, value)
	}
}
