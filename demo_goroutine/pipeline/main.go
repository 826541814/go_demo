package main

import (
	"fmt"
	"os"
	"bufio"
)
func mergeDemo(){
	a1:=ArraySource(3,2,6,7,3,1,3,5)
	a2:=ArraySource(9,8,7,6,5,4,3,2,1)
	// p := InMemSort(ArraySource(3,2,4,7,6,4,2,1,3,4,5,6,3))
	p1:=InMemSort(a1)
	p2:=InMemSort(a2)
	p:= Merge(p1,p2)
	// for {
	// 	if num,ok := <-p; ok {
	//		fmt.Println(num)
	//	}else{
	//		break
	//	}
	// }
	for v:=range p{
		fmt.Println(v)
	}
}

func old_main(){
	const filename = "large.in"
	const n=10000000
	file,err := os.Create(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	p:=RandomSource(n)
	writer := bufio.NewWriter(file)
	WriterSink(writer,p)
	writer.Flush()
	file,err = os.Open(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	p =ReaderSource(bufio.NewReader(file),-1)
	count :=0
	for v:= range p{
		fmt.Println(v)
		count++
		if count >=100{
			break
		}
	}
}

