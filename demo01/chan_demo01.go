package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		fmt.Println("timeSleep 5s")
		time.Sleep(time.Second * 5)
		<-ch
	}()
	ch <- 1
	fmt.Println("即将阻塞")
	ch <- 1
	fmt.Println("ch ")
}
