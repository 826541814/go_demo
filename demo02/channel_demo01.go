package main

import (
	"fmt"
	"time"
)

func toPrint(startCh <-chan string, stopCh <-chan string) {
	// for message := range stopCh {
	// 	fmt.Println(message)
	// }
	for {
		select {
		case start := <-startCh:
			fmt.Println(start)
		case stop := <-stopCh:
			fmt.Println(stop)
		default:
			fmt.Println("sleep")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	startCh := make(chan string, 10)
	stopCh := make(chan string, 10)
	go toPrint(startCh, stopCh)
	stopCh <- "aaaa"
	stopCh <- "aaaa"
	stopCh <- "aaaa"
	stopCh <- "aaaa"
	startCh <- "zzzz"
	startCh <- "zzzz"
	startCh <- "zzzz"
	startCh <- "zzzz"
	startCh <- "zzzz"
	startCh <- "zzzz"
	time.Sleep(time.Second * 3)
}
