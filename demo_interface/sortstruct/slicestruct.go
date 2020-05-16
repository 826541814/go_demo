package main
import (
	"fmt"
	"sort"
)
type HeroKind int

const(
	None = iota
	Tank
	Assassin
	Mage
)
type Hero struct{
	Name string
	Kind HeroKind
}
func main(){
	heros := []*Hero{
		{"lvbu",Tank},
		{"libd",Assassin},
		{"daji",Mage},
		{"diaochan",Assassin},
		{"guanyu",Tank},
		{"zhugeliang",Mage},
	}
	sort.Slice(heros,func(i,j int) bool{
		if heros[i].Kind != heros[j].Kind{
			return heros[i].Kind<heros[j].Kind
		}
		return heros[i].Name<heros[j].Name
	})
	for _,v := range heros{
		fmt.Printf("%+v\n",v)
	}
}
