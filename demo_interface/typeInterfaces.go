package main

import(
	"io"
)
type Socket struct{
}
func (s *Socket) Write(p []byte) (n int,err error){
	n=0
	return n,err
}
func (s *Socket) Close() error{
	return nil
}

type Writer interface{
	Write(p []byte) (n int,err error)
}
type Closer  interface{
	Close() error
}

func usingWriter(writer io.Writer){
	writer.Write(nil)
}
func usingCloser(closer io.Closer){
	closer.Close()
}
func main(){
	s:=new(Socket)
	usingCloser(s)
	usingWriter(s)
}
