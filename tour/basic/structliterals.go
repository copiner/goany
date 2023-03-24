package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{3,7}
	v2 = Vertex{X:9}
	v3 = Vertex{}
	p  = &Vertex{3,7}
)

func main(){

	fmt.Println(v1, p, v2, v3)
}
