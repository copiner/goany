package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}


func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}


func (v Vertex) Neet(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}


func main(){
	v := Vertex{3, 4}
	v.Neet(10)//with a value receiver, the Neet method operates on a copy of original Vertex value
	fmt.Println(v)
	v.Scale(10)//With a pointer receiver to change the Vertex value
	fmt.Println(v)//{30, 40}
	fmt.Println(v.Abs())
}
