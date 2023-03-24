/*
   type Tree struct {
       Left *Tree
       Value int
       Right *Tree
   }
*/
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int){
	if t.Left != nil{
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil{
		Walk(t.Right, ch)
	}
}



func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	Walk(t1, ch1)
	Walk(t2, ch2)

	for i := 0; i < 10; i++{
		v1 := <-ch1
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}

	return true
}

func main(){
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(3)))		
}
