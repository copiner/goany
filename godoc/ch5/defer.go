
package main

import "fmt"

func main(){

	r := kao()
	fmt.Println(r)
	s := nen()
	fmt.Println(s)

}

func kao()(r int){
	defer func(){
		r++
	}()
	return 1
}


func nen()(r int){

	defer func(r int){
		r++
	}(r)

	return 1
}
