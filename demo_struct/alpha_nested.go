package main
import "fmt"

type BasicColor struct{
	R,G,B float32
}
type Color struct{
	BasicColor
	G int32
	Alpha float32
}
func main(){
	color := &Color{}
	color.R=1
	color.G=1
	color.BasicColor.G=111
	color.B=0
	color.Alpha=1
	fmt.Println(*color)
	fmt.Printf("%+v\n",color)
}
