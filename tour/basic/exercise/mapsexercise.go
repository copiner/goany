package main

import (
	//"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WorldCount(s string) map[string]int{
	count := make(map[string]int)
	for _,v := range strings.Fields(s){
		count[v] = count[v] + 1
	}

	return count
}

func main(){
	wc.Test(WorldCount)
	//	a := "a b d g i p"
	//	fmt.Println(WorldCount(a))
}
