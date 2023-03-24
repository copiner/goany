package main

import "fmt"

type Vertex struct{
	Lat, Long float64
}

var m map[string]Vertex

func main(){
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		116.397128, 39.916527,
	}

	fmt.Println(m["Bell Labs"])
}

