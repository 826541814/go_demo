package main
import (
	"fmt"
	// "time"
	"runtime"
)

func main(){
	// var times = 4
	runtime.GOMAXPROCS(8)
	fmt.Println(runtime.NumCPU())
	go func(times int){
		for{
			fmt.Println("tick",times)
			// time.Sleep(time.Second)
			times++
		}
	}(5)
	var input string
	fmt.Scanln(&input)
}
