package main

import (
	"fmt"
	"time"
)

func main() {
	var result int
	go intervalSum(&result, 0, 50000)
	go intervalSum(&result, 50000, 100000)
	time.Sleep(time.Second) // give goroutines time to finish

	fmt.Println(result)

	otherResult := 0

	for i := 0; i < 100000; i++ {
		otherResult += i
	}
	fmt.Println(otherResult)
}

func intervalSum(destination *int, start, end int) {
	for i := start; i < end; i++ {
		*destination++
	}
}
