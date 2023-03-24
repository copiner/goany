
package main

import "fmt"

func main(){
	var a [5]int

	for i, v := range a{
		fmt.Printf("%d %d\n", i, v)
	}

	b := [...]int{9, 6, 3}
	for _, v := range b{
		fmt.Printf("%d\n", v)
	}

	c := [2]int{3,4}
	d := [...]int{3,4}

	fmt.Println(c == d)
}
