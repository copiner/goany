package main

import "fmt"

func fibonacci() func() int{
	one := 0
	two := 1
	return func() int{
		result := one
		one = two
		two = result + two
		return result
	}
}

func main(){
	f := fibonacci()
	for i := 0; i < 10; i++{
		fmt.Println(f())
	}
}
