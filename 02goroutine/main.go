package main

import (
	"fmt"
	"sync"
)

var c = make(chan bool, 10)
var wg = sync.WaitGroup{}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//fmt.Println(runtime.NumCPU())
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(i)
	}
	//for i := 0; i<10; i++ {
	//	<- c
	//}
	wg.Wait()
}

func Go(index int) {
	a := 1
	for i := 0; i < 100000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}
