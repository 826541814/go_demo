package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	//"time"
	"sync"
)

/*
1.channel就好像一个传送带，可以源源不断的往里面放数据，只要不close就可以永远发送数据。
2.如果channel里面没有数据，接收方会阻塞。
3.如果没有人正在等待channel的数据，发送方会阻塞。
4.从一个close的channel取数据永远不会阻塞，同时获取的是默认值
*/

func Print_url(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("11111", err)
	}
	defer res.Body.Close()
	fmt.Printf("网页的状态返回码是：%v\n", res.Status)

}

func work(resave chan string, wg *sync.WaitGroup) {

	for {
		url, ok := <-resave //如果channel里面没有数据，接收方会阻塞。用url,ok两个参数来接受channel，可以从channel会返回两个参数。url表示接受的数据。ok接受的信息是channel是否结束。
		//fmt.Println(url,ok)
		if !ok { //判断channel是否关闭。
			break
		}
		Print_url(url)
	}

	//for i := range resave { //range语法一值在接受数据，而channel可以源源不断的往里面添加数据。从一个close的channel取数据永远不会阻塞，同时获取的是默认值。
	//    Print_url(i)
	//}
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	//var num = 3
	//make_string := make(chan string)
	//for i := 0;i<num ;i++  {
	//    go work(make_string,wg)
	//}
	//wg.Add(num)
	fmt.Println(time.Now())
	make_string := make(chan string)
	for i := 0; i < 10; i++ {
		wg.Add(1) //启用的协程数和add的进程池保持一致。
		go work(make_string, wg)
	}

	urls := []string{"http://www.baidu.com", "http://www.qq.com", "http://www.weixin.com"}
	for _, i := range urls {
		for j := 0; j <= 20; j++ {
			make_string <- i //如果没有人正在等待channel的数据，发送方会阻塞。
		}

	}
	close(make_string) //channel就好像一个传送带，可以源源不断的往里面放数据，只要不close就可以永远发送数据。
	//time.Sleep(5 * time.Second)

	wg.Wait() //等待
	fmt.Println(time.Now())
}
