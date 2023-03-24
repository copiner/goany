package main

import(
	"fmt"
	"github.com/copiner/gogz/helper"
)

func main(){
	str := "hello world"
	fmt.Println(helper.Md5(str))
	helper.Help()
}
