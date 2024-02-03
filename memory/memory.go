package memory

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"unsafe"

	"sandboxProject/pkg"
)

type Memory struct{}

func (a *Memory) Start() {
	defer pkg.ExecuteTime("Start")()

	// Run server to get pprof data
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	finishChan := make(chan string)
	go exp1(finishChan)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case doneMsg := <-done:
		fmt.Println(doneMsg)
	case finishMsg := <-finishChan:
		fmt.Println(finishMsg)
	}
}

func exp1(finishChan chan string) {
	var Stri string = "123123123sadfasdfasdfw345434tertr"
	sizeOfInt := unsafe.Sizeof(Stri)
	fmt.Printf("Size of Stri: %d bytes, len: %d\n", sizeOfInt, len(Stri))
	finishChan <- PointerFoo(&Stri)
}

func PointerFoo(dfg *string) string {
	sizeOfInt := unsafe.Sizeof(dfg)
	return fmt.Sprintf("Size of Pointer: %d bytes\n", sizeOfInt)
}
