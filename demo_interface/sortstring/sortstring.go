package main

import (
	"fmt"
	"sort"
)
type MyStringList []string

func (m MyStringList) Len() int{
	return len(m)
}
func (m MyStringList) Less(i,j int) bool{
	return m[i]<m[j]
}
func (m MyStringList) Swap(i,j int){
	m[i],m[j] = m[j],m[i]
}
func main(){
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	slice :=[]string{ 
		"2. 2kill",
		"1. 1kill",
		"3. 3kill",
	}	
	sort.Sort(names)
	for _,v := range names {
		fmt.Printf("%s\n",v)
	}
	fmt.Println("__________________")
	sort.Strings(slice)
	for _,v := range slice{
		fmt.Printf("%s\n",v)
	}
	//fmt.Printf(sort.Sort(slice))
}
