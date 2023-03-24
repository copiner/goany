
package main

import "fmt"

func main(){

	bar()
}

func bar(){
	defer func(){
		fmt.Println("bar")
	}()

	foo()

	fmt.Println("bar end")
}

func foo(){

	defer func(){
		fmt.Println("foo")
	}()
	/*	defer func(){
		if err := recover(); err != nil{
			fmt.Println("recover")
		}
		fmt.Println("foo")
	}()*/

	panic("foo")
}

