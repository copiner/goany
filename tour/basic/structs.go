package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
func main(){
	fmt.Println(Vertex{3,7})

	v := Vertex{4,8}
	v.X = 9
	fmt.Println(v.X)
}
