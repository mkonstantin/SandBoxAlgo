package memory

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sandboxProject/pkg"
)

type MemoryTwo struct{}

const kingArrow = 1

func (a *MemoryTwo) Start() {
	defer pkg.ExecuteTime("Start")()

	// Run server to get pprof data
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	finishChan := make(chan string)
	//go handle(finishChan)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(1 * time.Second)

	counter := 0

	for {
		select {
		case doneMsg := <-done:
			fmt.Println("doneMsg:", doneMsg)
			ticker.Stop()
			return
		case finishMsg := <-finishChan:
			fmt.Println("finishMsg:", finishMsg)
			ticker.Stop()
			return
		case <-ticker.C:
			counter++
			go handle(counter)
		}
	}
}

func handle(number int) {
	fmt.Println("handle start", number)
	time.Sleep(3 * time.Second)
	continueAction(number)
	fmt.Println("handle stop", number)
}

func continueAction(number int) {
	fmt.Println("continueAction start", number)

	//m := runtime.MemStats
	//runtime.ReadMemStats(&m)

	time.Sleep(1 * time.Second)
	fmt.Println("const:", kingArrow)
	fmt.Println("continueAction stop", number)
}
