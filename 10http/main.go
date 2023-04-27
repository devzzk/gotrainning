package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	var success, failed int
	times := make([]time.Duration, 1000)
	for i := 0; i < 1000; i++ {
		start := time.Now()
		req, err := http.NewRequest("GET", "http://localhost:8000/api/v1/students", nil)
		times = append(times, time.Since(start))
		if err != nil {
			failed++
			//continue
		}
		resp, err := client.Do(req)
		if err != nil {
			failed++
			//continue
		}

		defer resp.Body.Close()
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			failed++
			//continue
		}
		success++
	}
	min := extremum(times, false)
	max := extremum(times, true)
	fmt.Printf("max: %v; min: %v", max, min)
	fmt.Printf("success rquest:%d;  failed request %d", success, failed)
}

func extremum(slice []time.Duration, isMax bool) any {

	en := slice[0]
	// 遍历切片，找到最大值
	for _, num := range slice {
		if isMax {
			if num > en {
				en = num
			}
		} else {
			if num < en {
				en = num
			}
		}

	}
	return en
}
