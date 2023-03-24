package main

import "fmt"

func main(){
	sum := 0
	for i:=0; i<10; i++{
		sum += i
	}

	fmt.Println(sum)

	total := 1
	for ; total < 1000; {
		total += total
	}

	fmt.Println(total)

	plus := 1
	for plus < 1000 {
		plus += plus
	}

	fmt.Println(plus)
}
