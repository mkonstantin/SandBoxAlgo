package concurency

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"sandboxProject/pkg"
)

type GoroutineOne struct{}

func (a *GoroutineOne) Start() {
	defer pkg.ExecuteTime("Start")()

	//GOMAXPROCSExp()
	//WaitGroupExp()

	ChannelsWithBufferExp()
}

func WaitGroupEx() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	simpleFunc := func(id int) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(3 * time.Second)
		fmt.Printf("Горутина %d закончила выполнение \n", id)
	}

	go simpleFunc(1)
	go simpleFunc(2)

	wg.Wait()
}

func GOMAXPROCSExp() {
	arr := []int{1, 2, 3, 4, 5}

	runtime.GOMAXPROCS(1)

	wg := sync.WaitGroup{}

	for _, item := range arr {
		wg.Add(1)

		go func(item int) {
			defer wg.Done()

			fmt.Println(item)
		}(item)
	}

	wg.Wait()
}

// Mutex
func ArrayExp1() {
	arr := []int{1, 2, 3, 4, 5}

	sdf := sync.Mutex{}

	sdf.Lock()

	sdf.Unlock()

	wg := sync.WaitGroup{}

	for _, item := range arr {
		wg.Add(1)

		go func(item int) {
			defer wg.Done()

			fmt.Println(item)
		}(item)
	}

	wg.Wait()
}

func WaitGroupExp() {
	number := 10000
	counter := 0

	// так как горутины выполняться много позже чем проход по циклу Фор и программа завершиться, нужно подождать когда они выполняться
	wg := sync.WaitGroup{}
	wg.Add(number)

	// чтобы при конкуренции блокировать запись чтобы не терять записи при конкуренции
	mu := sync.Mutex{}

	for i := 0; i < number; i++ {
		go func() {
			defer wg.Done()

			time.Sleep(time.Microsecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}

func RWMutexExp() {
	number := 10000
	counter := 0

	// так как горутины выполняться много позже чем проход по циклу Фор и программа завершиться, нужно подождать когда они выполняться
	wg := sync.WaitGroup{}
	wg.Add(number)

	// чтобы при конкуренции блокировать запись чтобы не терять записи при конкуренции
	//mu := sync.Mutex{}

	// блокировка на чтение
	rwmu := sync.RWMutex{}

	for i := 0; i < number/2; i++ {
		go func() {
			defer wg.Done()

			rwmu.Lock()

			time.Sleep(time.Microsecond)
			counter++

			rwmu.Unlock()
		}()

		go func() {
			defer wg.Done()

			rwmu.RLock()

			time.Sleep(time.Microsecond)
			_ = counter

			rwmu.RUnlock()
		}()
	}

	wg.Wait()

	fmt.Println("counter %w", counter)
}

func ChannelsWithBufferExp() {
	number := 10
	counter := 0

	// так как горутины выполняться много позже чем проход по циклу Фор и программа завершиться, нужно подождать когда они выполняться
	// для этого юзается WaitGroup
	wg := sync.WaitGroup{}

	mu := &sync.Mutex{}

	//rmu := &sync.RWMutex{}

	// Создаем канал с буффером = number
	chl := make(chan int, number) // - with Buffer

	for i := 0; i < number; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Microsecond)

			chl <- counter
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait() // ожидаем когда отработают первые горутины

	for i := 0; i < number; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Microsecond)

			fmt.Println(<-chl)
			mu.Unlock()
		}()
	}

	wg.Wait() // ожидаем когда отработают вторые горутины

	fmt.Println("Done")
}
