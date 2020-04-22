package main

import (
	"fmt"
	"time"
)

// 数据生产者
func producter(header string, channel chan<- string, stopCh chan string) {
	for i := 1; i <= 5; i++ {
		channel <- fmt.Sprintf("%s:%v", header, i)
		// time.Sleep(time.Second)
	}
	for i := 1; i <= 5; i++ {
		fmt.Println("write stopCh")
		// stopCh <- fmt.Sprintf("stopCh:%v", i)
		// time.Sleep(time.Second)
		stopCh <- "aaa"
		fmt.Sprintf("get %s", <-stopCh)
	}
	// fmt.Println(<-stopCh)
	// stopCh <- header
	fmt.Println("输入完成")

}

// 数据消费者
func consumer(channel <-chan string, stopCh <-chan string) {
	// for message := range channel {
	// 	// fmt.Println("回车继续")
	// 	// fmt.Scanln()
	// 	// message := <-channel
	// 	fmt.Println(message)
	// }
	// i := 0
	// var header string
	// for {
	// 	message := <-channel
	// 	_= <-stopCh
	// 	fmt.Sprintf("%s 和 %s", message, stop)
	// }
	for {
		select {
		case message := <-channel:
			fmt.Println(message)
		case header := <-stopCh:
			// break
			fmt.Sprintf("%s finished", header)
		// time.Sleep(time.Second)
		default:
			time.Sleep(time.Second)
			fmt.Println("sleep....")
		}
	}
	// end:
	// 	fmt.Println("finished")

}

func main() {
	channel := make(chan string)
	stopCh := make(chan string, 10)
	fmt.Println("start")
	consumers := consumer
	// consumers := new func(interface)
	// consumers = consumer
	go producter("cat", channel, stopCh)
	go producter("dog", channel, stopCh)
	go producter("zhangsan", channel, stopCh)
	consumers(channel, stopCh)
	time.Sleep(time.Second * 5)
	fmt.Println("finished")
}
