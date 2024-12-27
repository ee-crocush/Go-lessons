package main

import (
	"fmt"
	"os"
	"time"
)

func Example() {
	fmt.Println("we start")

	go func() {
		for i := 0; i < 15; i++ {
			time.Sleep(time.Microsecond) // Имитируем некие вычисления
			fmt.Println("Counter is", i)
		}
	}()

	go func() {
		for i := 0; i < 15; i++ {
			time.Sleep(time.Microsecond) // Имитируем некие вычисления
			fmt.Println("Other counter is", i)
		}
	}()

	for i := 0; i < 15; i++ {
		fmt.Println("Main counter is", i)
	}

	fmt.Println("we finished")

	time.Sleep(time.Second) // Даем время проинициализироваться и выполнить свой код
}

func main() {
	var n int
	fmt.Print("Введите число: ")

	// Сканируем ввод
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		return
	}

	fmt.Printf("Вы ввели число: %d\n", n)

	// Запускаем n горутин
	for i := 0; i < n; i++ {
		i := i
		go func() {
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
				time.Sleep(1 * time.Second)
			}
		}() // Передаём значение `i` как параметр, чтобы избежать проблемы замыкания
	}
	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)

	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		return
	}
}
