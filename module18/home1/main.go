/*
Создайте программу, которая запускает 5 рутин, каждая из которых
печатает свой порядковый номер 10 раз. Добиться такого результата
за один вызов wg.Add.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	const gorutineCount = 5
	const iterations = 10

	wg.Add(gorutineCount)

	for i := 0; i < gorutineCount; i++ {
		go func(i int) {
			for j := 0; j < iterations; j++ {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	time.Sleep(time.Second)
}
