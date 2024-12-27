package main

import (
	"fmt"
	"sync"
	"time"
)

var mu = sync.Mutex{}

func main() {
	var n int
	fmt.Print("Введите число: ")

	// Сканируем ввод
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(fmt.Errorf("ошибка: %v", err))
		return
	}

	result := 0
	// Горутина для увеличения значения каждую секунду
	go func() {
		for {
			time.Sleep(time.Second)
			mu.Lock()
			result++
			mu.Unlock()
		}
	}()

	// Горутина для вывода значения каждые ⅕ секунды
	go func() {
		for {
			mu.Lock()
			fmt.Printf("Result: %d\n", result)
			mu.Unlock()
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second * time.Duration(n))
}
