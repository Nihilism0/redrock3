package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var ch1 = make(chan int, 1)

func main() {
	wg1.Add(2)
	go func() {
		for i := 0; i < 50; i++ {
			time.Sleep(time.Millisecond * 5) //等待,避免运行过快发生资源抢夺
			ch1 <- 1                         //堵塞,保证有序
			fmt.Println(i*2 + 1)

		}
		wg1.Done()
	}()
	go func() {
		for i := 1; i <= 50; i++ {
			time.Sleep(time.Millisecond * 5)
			<-ch1 //释放,保证有序
			fmt.Println(i * 2)

		}
		wg1.Done()
	}()
	wg1.Wait()
}
