package examples

import (
	"context"
	"fmt"
	"sync"
	"time"

	"sandboxProject/pkg"
)

type ExamplesTwo struct{}

func (a *ExamplesTwo) Start() {
	defer pkg.ExecuteTime("Start")()

	ch := make(chan string)
	pool := NewPool(ch, 5)

	fmt.Println("Start")
	pool.Start(context.Background())
	time.Sleep(1 * time.Second)

	fmt.Println("Send data")
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("string %d", i)
	}
	time.Sleep(3 * time.Second)

	fmt.Println("Stop")
	pool.Stop()
}

type WorkerPool interface {
	Start(ctx context.Context)
	Stop()
}

type Pool struct {
	done     chan int
	input    chan string
	wrkCount int
}

func NewPool(input chan string, wrkCount int) *Pool {
	return &Pool{
		done:     make(chan int),
		input:    input,
		wrkCount: wrkCount,
	}
}

func (p *Pool) Start(ctx context.Context) {
	wg := sync.WaitGroup{}
	wg.Add(p.wrkCount)

	for i := 0; i < p.wrkCount; i++ {
		go func(workerNumber int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("ctx.Done()", workerNumber)
					return
				case <-p.done:
					fmt.Println("p.done", workerNumber)
					return
				case <-p.input:
					fmt.Println(<-p.input, workerNumber)
				}
			}
		}(i)
	}
	wg.Wait()
}

func (p *Pool) Stop() {
	close(p.done)
}
