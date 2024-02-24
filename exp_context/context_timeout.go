package exp_context

import (
	"context"
	"fmt"
	"time"
)

type TimeoutContext struct{}

func (a *TimeoutContext) Start() {
	timeoutExp := NewTimeoutExp()
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	go timeoutExp.start(ctx)
	time.Sleep(3 * time.Second)
	fmt.Println("End", time.Now())
}

type TimeoutExp struct {
}

func NewTimeoutExp() *TimeoutExp {
	return &TimeoutExp{}
}

func (a *TimeoutExp) start(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("After", time.Now())
		return
	case <-ctx.Done():
		fmt.Println("Done", time.Now())
		return
	}
}
