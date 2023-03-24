
package main

import "fmt"

func appendInt(x []int, y int) []int {

	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x){

		z = x[:zlen]

	} else {
		
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
		
	}
	z[len(x)] = y

	return z
}


func appendIntm(x []int, y ...int) []int {

	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x){

		z = x[:zlen]

	} else {
		
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
		
	}
	copy(z[len(x):], y)

	return z
}


func main(){
	var x, y []int

	for i := 0; i < 10; i++{
		//y = appendInt(x, i)
		y = appendIntm(x, i, 9)		
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	
	var a []int

	a = append(a, 1)
	a = append(a, 3, 5)
	a = append(a, 7, 9)

	a = append(a, a...)	
	fmt.Println(a)
}
