package main

import (
	"fmt"
	"sync"
)

// Число сообщений от каждого источника
const messagesAmountPerGoroutine int = 5

// startDataSource создаёт источник данных.
// Он отправляет сообщения в канал и затем закрывает его
func startDataSource(start int) chan int {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := start; i < start+messagesAmountPerGoroutine; i++ {
			c <- i
		}
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}

// multiplexingFunc объединяет несколько каналов в один
func multiplexingFunc(channels ...chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedChan := make(chan int)

	// Функция для обработки каждого канала
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			multiplexedChan <- i
		}
	}

	// Добавляем количество каналов в счётчик
	wg.Add(len(channels))

	// Запускаем обработку каждого канала в отдельной горутине
	for _, c := range channels {
		go multiplex(c)
	}

	// Закрываем общий канал, когда все завершили работу
	go func() {
		wg.Wait()
		close(multiplexedChan)
	}()

	return multiplexedChan
}

func main() {
	// Создаём источники данных
	var dataSourceChans []chan int
	for i := messagesAmountPerGoroutine; i <= 20; i += messagesAmountPerGoroutine {
		dataSourceChans = append(dataSourceChans, startDataSource(i))
	}

	// Уплотняем каналы
	c := multiplexingFunc(dataSourceChans...)

	// Читаем данные из общего канала
	for data := range c {
		fmt.Println(data)
	}
}
