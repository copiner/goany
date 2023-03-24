//greatest common divisor

package main

import "fmt"

func gcd(x, y int) int{
	for y != 0{
		x, y = y, x%y
	}
	return x
}

func fib(n int) int{
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main(){
	g := gcd(18, 6)
	fmt.Println(g)
}
