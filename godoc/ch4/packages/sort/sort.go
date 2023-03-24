
package main

import (
	"fmt"
	"local/treesort"
)

func main(){

	arr := []int{2,10,8,4,5,3,7,9}

	treesort.Sort(arr[0:8])
	fmt.Println(arr)
}
