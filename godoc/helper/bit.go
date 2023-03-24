/*
   反码，补码，原码
   
   & 两位全为1，结果为1 否则为0
   | 两位有一个为1，结果为1，否则为0
   ^ 两位一个为0，一个为1，结果为1，否则为0

   >> 右移运算符：低位溢出，符号为不变，并且符号为补溢出的高位
   << 左移运算符：符号为不变，低位补0
*/

package main

import "fmt"

func main(){

	//2 补码 00000010
	//3 补码 00000011
	fmt.Println(2&3)

	fmt.Println(2|3)

	fmt.Println(2^3)

	//2  补码   00000010
	//-2 补码   1111 1110
	fmt.Println(-2^2)
	
	//1      0000 0001 -> 0000 0000 = 0
	a := 1>>2

	//1      0000 0001 -> 0000 0100 = 4
	b := 1<<2
	
	//-1     1000 0001 -> 0000 0001 = -1
	c := -1>>2
	
	//-1     1000 0001 -> 0000 0100 = -4
	d := -1<<2

	fmt.Println(a,b,c,d)

	
}
