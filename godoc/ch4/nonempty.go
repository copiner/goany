
package main

import "fmt"

/*
   nonempty returns a slice holding only the non-empty strings
   The underlying array is modified during the call
*/

func nonempty(strings []string) []string{

	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func nonempty2(strings []string) []string{

	out := strings[:0] //zero-length slice of original
	
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func main(){
	data := []string{"three", "", "four"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	data2 := []string{"three", "", "four"}
	fmt.Printf("%q\n", nonempty2(data2))
	fmt.Printf("%q\n", data2)	
}

