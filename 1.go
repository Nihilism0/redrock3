package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int, 1) //利用channel可以被堵住♂的特性
var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 50000; i++ {
		ch <- 1 //堵住,其他的兄弟们拿不了
		x = x + 1
		<-ch //释放,其他的兄弟们可以开抢了
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
