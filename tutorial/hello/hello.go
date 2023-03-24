/*
   edit the tutorial/hello module to use local tutorial/greetings
   go mod edit -replace tutorial/greetings=../greetings

   go mod tidy
*/
package main

import (
	"fmt"
	//"rsc.io/quote"
	"tutorial/greetings"
	"log"
)


func main() {
	
	//message := greetings.Hello("copiner")

	//fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	//message, err := greetings.Hello("Copiner")
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	
	if err != nil {
		log.Fatal(err)
	}

	
	//fmt.Println(message)
	fmt.Println(messages)		
}
