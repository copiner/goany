
package main

import "fmt"

const (
	a = 3
	b
	c = 7
	d
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main(){
	fmt.Println(a, b, c, d)
	fmt.Println(Sunday, Friday)	
}
