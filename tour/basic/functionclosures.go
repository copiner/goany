/*
   Go function may be closures.
   A closure is a function value that references varibales from outside its body
*/
package main

import "fmt"

func adder() func(int) int{//func(int)   int
	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}


func main(){
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++{
		fmt.Println(pos(i), neg(-2*i),)
	}
}
