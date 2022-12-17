package types

import (
	"fmt"
	"reflect"
)

type TypesExamples struct{}

func (a *TypesExamples) Start() {
	lessonOne()
}

func lessonOne() {
	// Посмотрим какой тип возвращает:
	Print(reflect.TypeOf(42))
	Print(reflect.TypeOf(anonymFunction()))

	// Если просто вывести то возвращает адрес памяти 0x10ad240
	Print(anonymFunction())
	// Чтобы получить значение, надо его определить на переменную
	var value1 = anonymFunction()
	Print(value1())

	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 5}
	fmt.Println(lesson(arr1, arr2))
}

//Преимуществом анонимных функций является то, что они имеют доступ к окружению, в котором они вызываются
func anonymFunction() func() string {
	str := "own string"
	return func() string {
		return fmt.Sprint(str)
	}
}

// type Option func()

func Print(value interface{}) {
	fmt.Println(value)
}

func lesson(arr1, arr2 []int) []int {
	var outArr []int
	for _, value := range arr1 {
		if contains(arr2, value) == false {
			outArr = append(outArr, value)
		}
	}
	return outArr
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
