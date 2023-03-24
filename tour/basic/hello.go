package main

import (
       "fmt"
       "time"
       "math"
       "math/rand"
)

func add(x int, y int) int{
     return x + y
}

func plus(x, y int) int{
     return x + y
}


func main(){
     fmt.Println("hello, world")
     fmt.Println("The time is", time.Now())
     fmt.Println("My favorite number is", rand.Intn(10))
     fmt.Println("My favorite number is", math.Sqrt(7))
     fmt.Println("My favorite number is", math.Pi)
     fmt.Println("My favorite number is", add(2,3))
     fmt.Println("My favorite number is", plus(2,3))     	
}
