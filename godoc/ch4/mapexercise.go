
package main

import "fmt"

//map equal
func equal(x, y map[string]int) bool{
	if len(x) != len(y){
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func main(){

	//ages := make(map[string]int)

	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}

	fmt.Println(ages["alice"])
	fmt.Println(ages["ali"])	

	var bloods map[string]int

	fmt.Println(bloods == nil)
	fmt.Println(len(bloods) == 0)	
}
