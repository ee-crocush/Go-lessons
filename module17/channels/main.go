package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c <- "Сообщение"
		fmt.Println("Горутина отработала!")
	}()
	//time.Sleep(100 * time.Millisecond)
	fmt.Println(<-c)
	wg.Wait()
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2 // Данные остаются в канале до их извлечения
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}
}
