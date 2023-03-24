//Newton's method

package main

import (
	"fmt"
	//"math"
)

func Sqrt(x float64) float64{
	z := float64(1)
	for i := 0; i < 10; i++{
		z -= (z*z - x) / (2*z)
	}
	return z
	//return math.Sqrt(x)
}

func main(){
	fmt.Println(Sqrt(2))
}
