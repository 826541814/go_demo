package main
import "fmt"
func helloWorld(id int,ch chan string){
	for {
		ch <- fmt.Sprintf("Hello world from"+"goroutine %d\n",id)
	}
}
func main(){
	ch := make(chan string)
	for i:=0; i<=500 ;i++{
		go helloWorld(i,ch)
	}
	for{
		msg:=<-ch
		fmt.Println(msg)
	}
}
