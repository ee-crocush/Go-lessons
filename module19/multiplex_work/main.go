package main

import (
	"fmt"
	"sync"
)

// Число сообщений от каждого источника
const messagesAmountPerGoroutine int = 5

// multiplexingFunc объединяет несколько каналов в один
func multiplexingFunc(done <-chan int, channels ...chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedChan := make(chan int)

	// Функция для обработки каждого канала
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for {
			select {
			case i := <-c:
				multiplexedChan <- i
			case <-done:
				return
			}
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

// Функция разуплотнения каналов
func demultiplexingFunc(dataSourceChan chan int, amount int) ([]chan int, <-chan int) {
	// Канал, сигнализатор выполнения
	var done = make(chan int)
	// Создаем слайс каналов
	var output = make([]chan int, amount)
	// Проходимся по каждому элементу слайса и инициализируем его как канал
	for i := range output {
		output[i] = make(chan int)
	}
	// Запускаем внешнюю горутину - отвечает за управление процессом разуплотнения данных.
	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		// запускаем горутину, которая читает данные из входного канала dataSourceChan
		go func() {
			defer wg.Done()
			// проходит по каждому элементу входного канала
			for v := range dataSourceChan {
				//передаёт элемент из входного канала во все каналы output.
				for _, c := range output {
					c <- v
				}
			}
		}()
		wg.Wait()
		// закрываем каналы
		close(done)
	}()
	return output, done
}

// startDataSource создаёт источник данных.
// Он отправляет сообщения в канал и затем закрывает его
func startDataSource() chan int {
	c := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= messagesAmountPerGoroutine; i++ {
				c <- i
			}
		}()
		wg.Wait()
		close(c)
	}()
	return c
}

func main() {
	// Создаём источник данных
	dataSourceChans := startDataSource()
	// Запускаем источник данных и уплотняем каналы
	consumers, done := demultiplexingFunc(dataSourceChans, 5)
	c := multiplexingFunc(done, consumers...)

	for data := range c {
		fmt.Println(data)
	}
}
