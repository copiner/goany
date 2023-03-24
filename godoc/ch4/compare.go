
package main

import "fmt"

//slice compare
func equal(x, y []string) bool{
	if len(x) != len(y){
		return false
	}

	for i := range x {
		if x[i] != y[i]{
			return false
		}
	}

	return true
}

func main(){
	var a []int
	a = nil
	
	b := []int(nil)
	c := []int{}

	d := make([]int, 0, 0)

	fmt.Println(len(a))
	if a == nil {
		fmt.Println("equal")
	}
	
	fmt.Println(len(b))	
	if b == nil {
		fmt.Println("equal")
	}

	fmt.Println(len(c))
	if c != nil {
		fmt.Println("other")
	}


	fmt.Println(len(d))
	if d != nil {
		fmt.Println("equal")
	}
	
}
