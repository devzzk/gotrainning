package main

import "fmt"

func main() {
	//固定长度数组，不可增加长度， 不可删减长度
	var arr = [3]int{1, 2, 3}

	//可变长度数组， 又称 slice (切片)
	var slic = []int{1, 2}

	fmt.Println(arr, slic)
}
