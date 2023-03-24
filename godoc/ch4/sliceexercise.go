
package main

import "fmt"

//remove
func remove(slice []int, i int) []int{
	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1]
}

func main(){
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))

	//stack
	stack := []int{}
	
	stack = append(stack, 7)
	stack = append(stack, 9)

	top := stack[len(stack)-1] //top of stack

	stack = stack[:len(stack)-1]	//pop

	fmt.Println(top)	
	fmt.Println(stack)
}
