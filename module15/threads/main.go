package main

import (
	"fmt"
	"sync"
	"time"
)

func SomeGoroutine() {
	var wg sync.WaitGroup
	num := 32
	for i := 0; i < num; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer func() {
				fmt.Printf("Goroutine %d finished\n", i+1)
				wg.Done()
			}()
			var c uint64
			for j := uint64(0); j < 1<<num; j++ {
				c++
			}
		}()
	}
	wg.Wait()
}

func HelloGoroutine() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(fmt.Sprintf("Hello from goroutine %d", i))
			time.Sleep(time.Second)
		}()
	}
	time.Sleep(time.Second)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to the channel
}

func main() {
	HelloGoroutine()
	fmt.Println("Main thread finished")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)], c)
	//go sum(s[len(s)/2:], c)

	fmt.Println(<-c)
}
