
package main

import "fmt"

func reverse(s []int){ //slice
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1{
		s[i], s[j] = s[j], s[i]
	}
}

func main(){
	a := [...]int{9, 8, 7, 6, 5} //array
	reverse(a[:])
	fmt.Println(a)

	s := []int{9, 8, 7, 6, 5} //slice

	reverse(s[:2])
	reverse(s[2:])

	reverse(s)
	fmt.Println(s)	
}
