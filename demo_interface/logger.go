package main

type LogWriter interface{
	Write(data interface{}) error
}
type Logger struct{
	writerList []LogWriter
}
func(l *Logger) RegisterWriter(writer LogWriter){
	l.writerList = append(l.writerList,writer)
}
func (l*Logger) Log(data interface{}){
	for _,writer := range l.writerList{
		writer.Write(data)
	}
}
func NewLogger() *Logger{
	return &Logger{}
}
type Screen struct{
	size int
}
func (s *Screen) Write(data interface{}) error{
	return nil
}
func main(){
	logger := NewLogger()
	screen := &Screen{}
	var write LogWriter
	write = screen
	logger.RegisterWriter(write)
	
}
