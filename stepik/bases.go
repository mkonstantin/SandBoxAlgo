package stepik

import (
	"fmt"
	"math"
	"time"
)

type Bases struct{}

func (s *Bases) Start() {
	welcome()
	fmt.Println(evclid())
	fmt.Println(resultTimes("Hello", 3))
	fmt.Println(language("en"))
}

// Приветствие, перевод времени в минуты
func welcome() {
	str := "1h30m"

	d, _ := time.ParseDuration(str)

	fmt.Println(fmt.Sprintf("%d min", int(d.Minutes())))
}

// рассчитать расстояние между двумя точками по эвклидовой формуле
func evclid() float64 {
	var x1, y1, x2, y2 float64
	x1 = 123.23
	x2 = 62.123
	y1 = 12.23
	y2 = 43.123
	d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return d
}

// Повторить строку source times раз
func resultTimes(source string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result = result + source
	}
	return result
}

// switch cases example
func language(code string) string {
	switch code {
	case "en":
		return "English"
	case "fr":
		return "French"
	case "ru":
		return "Russian"
	case "rus":
		return "Russian"
	default:
		return "Unknown"
	}
}
