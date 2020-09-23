package day01

import (
	"fmt"
	"time"
)

var jobs = make(chan int, 10)
var results = make(chan int, 10)

func woker(id int, job <-chan int, results chan<- int) {
	for j := range job {
		fmt.Printf("wokerid:%d, start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("wokerid:%d, end job:%d\n", id, j)
		results <- j * 2
	}
}

/* channel */
func Nalademp() {
	/*开3个协程 */
	for i := 1; i <= 3; i++ {
		go woker(i, jobs, results)
	}
	/* 开5个任务 */
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	//read results
	for a := 1; a <= 5; a++ {
		res := <-results
		fmt.Printf("results:%d\n", res)
	}
}
