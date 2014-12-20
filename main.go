package main

import (
	"fmt"
	"runtime"
	"sync"
)

func (t *Thread) test(text string) {
	t.Wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%s %d\n", text, i)
		}
		t.Wg.Done()
	}()

}

type Thread struct {
	Wg *sync.WaitGroup
}

func main() {
	t := &Thread{}

	runtime.GOMAXPROCS(runtime.NumCPU())
	t.Wg = new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		text := fmt.Sprintf("go routine %d", i)
		t.test(text)
	}

	t.Wg.Wait()

	fmt.Println("End")
}
