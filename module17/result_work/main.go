package main

import (
	"fmt"
)

func gotoutine1() <-chan int {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	return c
}

func gotoutine2() <-chan int {
	c := make(chan int)
	go func() {
		c <- 2
		select {}
	}()
	return c
}

func main() {
	fmt.Printf("Отработала горутина №%d\n", <-gotoutine1())
	fmt.Printf("Отработала горутина №%d\n", <-gotoutine2())
	fmt.Println("Выполнено")
}
