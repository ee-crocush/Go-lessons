package main

import (
	"fmt"
	"semaphore/semaphore" // Путь к пакету semaphore
	"sync"
	"time"
)

func main() {
	// Создаём семафор с таймаутом 1 секунда
	sem := semaphore.NewSemaphore(1 * time.Second)

	var wg sync.WaitGroup

	// Функция для работы с ресурсом, используя семафор
	worker := func(id int) {
		defer wg.Done()
		fmt.Printf("Worker %d: пытается захватить семафор\n", id)
		if err := sem.Acquire(); err != nil {
			fmt.Printf("Worker %d: ошибка захвата семафора: %s\n", id, err)
			return
		}
		fmt.Printf("Worker %d: захватил семафор\n", id)

		// Имитация работы с ресурсом
		time.Sleep(500 * time.Millisecond)

		// Освобождение семафора
		if err := sem.Release(); err != nil {
			fmt.Printf("Worker %d: ошибка освобождения семафора: %s\n", id, err)
		} else {
			fmt.Printf("Worker %d: освободил семафор\n", id)
		}
	}

	// Запуск 5 горутин
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("Все работы завершены")
}
