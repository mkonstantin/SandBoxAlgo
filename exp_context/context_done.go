package exp_context

import (
	"context"
	"fmt"
	"time"
)

type DoneContext struct{}

func (a *DoneContext) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	algo := NewAlgo(cancel)
	go algo.Start(ctx)

	time.Sleep(1 * time.Second)
	algo.Stop()
	time.Sleep(1 * time.Second)
}

type Algo struct {
	cancelFunc context.CancelFunc
}

func NewAlgo(cancelFunc context.CancelFunc) *Algo {
	return &Algo{
		cancelFunc: cancelFunc,
	}
}

func (a *Algo) Start(ctx context.Context) {
	input := make(chan int, 5)
	input <- 1
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("After")
		return
	case <-ctx.Done():
		fmt.Println("Done")
		return
	}
}

func (a *Algo) Stop() {
	fmt.Println("Stopping")
	a.cancelFunc()
}
