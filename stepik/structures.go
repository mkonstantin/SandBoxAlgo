package stepik

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type StructuresOne struct {
}

func (s *StructuresOne) Start() {
	password := password{
		value:     "12312312",
		validator: and(minlen(8), digits, letters),
	}

	fmt.Println(password.isValid())
}

// validator проверяет строку на соответствие некоторому условию
// и возвращает результат проверки
type validator func(s string) bool

// digits возвращает true, если s содержит хотя бы одну цифру
// согласно unicode.IsDigit(), иначе false
func digits(s string) bool {
	for _, i2 := range s {
		if unicode.IsDigit(i2) {
			return true
		}
	}
	return false
}

// letters возвращает true, если s содержит хотя бы одну букву
// согласно unicode.IsLetter(), иначе false
func letters(s string) bool {
	for _, i2 := range s {
		if unicode.IsLetter(i2) {
			return true
		}
	}
	return false
}

// minlen возвращает валидатор, который проверяет, что длина
// строки согласно utf8.RuneCountInString() - не меньше указанной
func minlen(length int) validator {
	return func(s string) bool {
		return utf8.RuneCountInString(s) >= length
	}
}

// and возвращает валидатор, который проверяет, что все
// переданные ему валидаторы вернули true
func and(funcs ...validator) validator {
	if len(funcs) == 0 {
		return nil
	}
	return func(s string) bool {
		for _, f := range funcs {
			if !f(s) {
				return false
			}
		}
		return true
	}
}

// or возвращает валидатор, который проверяет, что хотя бы один
// переданный ему валидатор вернул true
func or(funcs ...validator) validator {
	if len(funcs) == 0 {
		return nil
	}
	return func(s string) bool {
		for _, f := range funcs {
			if f(s) {
				return true
			}
		}
		return false
	}
}

// password содержит строку со значением пароля и валидатор
type password struct {
	value string
	validator
}

// isValid() проверяет, что пароль корректный, согласно
// заданному для пароля валидатору
func (p *password) isValid() bool {
	return p.validator(p.value)
}
