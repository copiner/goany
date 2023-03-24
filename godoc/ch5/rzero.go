
package main

import "fmt"

func main(){
	
	r := rzero(9)
	fmt.Println(r)
	
}

func rzero(a int) (r int){

	defer func(){
		if p := recover(); p != nil {
			r = a
		}
	}()

	panic(a)
}
