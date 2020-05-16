package main

import (
	"fmt"
	"errors"
	"os"
)
type fileWriter struct{
	file *os.File
}
func (f *fileWriter) SetFile(filename string) (err error){
	if f.file!=nil{
		f.file.Close()
	}
	f.file,err = os.Create(filename)
	return err
}
func (f *fileWriter) Write (data string) error{
	if f.file == nil{
		return errors.New("file not created")
	}
	str := fmt.Sprintf("%v\n",data)
	_,err := f.file.Write([]byte(str))
	return err
}
func newFileWriter() *fileWriter{
	return &fileWriter{}
}
func main(){
	writer := newFileWriter()
	writer.SetFile("createFile.test")
	value := "hello world"
	writer.Write(value)
	writer.file.Close()
}
