package main

import (
	"os"
	"bufio"
)

func createPipeline(filename string,fileSize,chunkCount int) <-chan int{
	chunkSize=fileSize/chunkCount
	sortResults:=[]<-chan int{}
	for i:=0; i<chunkCount; i++{
		file,err := os.Open(filename)
		if err != nil{
			panic(err)
		}
		file.Seek(int64(i*chunkSize),0)
		source:=ReaderSource(bufio.NewReader(file),chunkSize)
		sortResults=append(sortResults,InMemSort(source))
	}
	return MergeN(sortResults...)
}

func main(){
	p:=createPipeline("small.in",400,4)
	writeToFile(p)
	printFile()
}
