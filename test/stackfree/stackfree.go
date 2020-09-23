package main

import (
	"fmt"
	"sync"
)

var res int
var wait sync.WaitGroup
var mux sync.Mutex
var resCh = make(chan int, 1)

func add12(i int) {
	mux.Lock()
	res += i
	mux.Unlock()
	wait.Done()
}

func main12() {
	for i := 0; i <= 100; i++ {
		wait.Add(1)
		go add12(i)
	}
	wait.Wait()
	fmt.Println(res)
}
func main() {

	for i := 1; i <= 100000; i++ {
		res += i
	}
	resCh <- res
	fmt.Println(<-resCh)

}
