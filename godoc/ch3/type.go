
package main

import "fmt"



func main(){

	//int8 -128 ~ 127, uint8 0 ~ 255
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	
	modals := []string{"gold", "silver", "bronze"}

	for i := len(modals) - 1; i >= 0; i-- {
		fmt.Println(modals[i])
	}

	//float32 float64
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

}



