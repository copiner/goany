package main

import "fmt"

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main(){
	var s []int
	printSlice(s)

	s = append(s,9)
	printSlice(s)

	s = append(s, 7)
	printSlice(s)

	s = append(s, 3, 2, 1)
	printSlice(s) //length 5  capacity 6
}
