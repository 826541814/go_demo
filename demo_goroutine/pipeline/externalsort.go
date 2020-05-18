package main
import (
	"fmt"
	"os"
	"bufio"

)

func createPipeline(filename string,fileSize,chunkCount int)<-chan int{
	chunkSize:=fileSize/chunkCount
	sortResults:=[]<-chan int{}
	Init()
	for i:=0; i<chunkCount; i++{
		file,err:=os.Open(filename)
		if err!=nil{
			panic(err)
		}
		file.Seek(int64(i*chunkSize),0)
		source :=ReaderSource(bufio.NewReader(file),chunkSize)
		sortResults=append(sortResults,InMemSort(source))
	}
	return MergeN(sortResults...)
}

func writeToFile(p<-chan int,filename string){
	file,err := os.Create(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	writer :=bufio.NewWriter(file)
	defer writer.Flush()

	WriterSink(writer,p)
}

func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	p:=ReaderSource(file,-1)
	for v:=range p{
		fmt.Println(v)
	}

}

func main(){
	p:=createPipeline("large.in",80000000,8)
	writeToFile(p,"large.out")
	printFile("large.out")
}
