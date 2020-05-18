package main

import (
	"fmt"
	"time"
)

func running(mult int){
	var times int
	for {
		times++
		fmt.Println("tick",times)
		time.Sleep(time.Second*1)
	}
}

func main() {
	go running(1)
	var input string
	fmt.Scanln(&input)

}
