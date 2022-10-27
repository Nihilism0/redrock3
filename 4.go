package main

import (
	"fmt"
	"time"
)

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ { //问题出在所有人同时拿走i,还没等print,最后的9就覆盖了前面的所有操作,这时候,需要让其慢一点
		time.Sleep(time.Millisecond * 10)
		if i == 9 { //换个位置
			over <- true
		}
		go func() {
			fmt.Println(i)
		}()

	}
	<-over
	fmt.Println("over!!!")
}
