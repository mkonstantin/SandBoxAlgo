package pkg

import (
	"fmt"
	"time"
)

func ExecuteTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s execution time: %v\n", name, time.Since(start))
	}
}
