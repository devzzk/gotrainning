package main

import "fmt"

func main() {
	//byte 占一个字节
	ar := [3]byte{'a', 'b', 'c'}

	fmt.Println(ar) // '97' '98' '99' ASCII编码顺序
}
