package main

import "fmt"

func pslice(x []int){
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}

func main(){
	s := make([]int, 3, 6)
	pslice(s)

	s = append(s, 2)
	pslice(s)

	s = append(s, 5)
	s = append(s, 7)
	s = append(s, 8) //GC
	pslice(s)

	fmt.Println("------")
	a := make([]int, 3, 6)
	b := a[1:3]

	pslice(a)
	pslice(b)

	b = append(b,9)

	pslice(a)
	pslice(b)

	b = append(b,4)
	b = append(b,3)
	b = append(b,1)

	pslice(a)
	pslice(b)

	c := a[:6]
	pslice(c)
}
