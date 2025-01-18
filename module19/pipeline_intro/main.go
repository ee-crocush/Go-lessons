package main

import "fmt"

//func main() {
//	// Стадия умножения потока целых чисел на заданное число
//	multiply := func(values []int, multiplier int) []int {
//		multipliedValues := make([]int, len(values))
//		for i, v := range values {
//			multipliedValues[i] = v * multiplier
//		}
//		return multipliedValues
//	}
//	// Стадия суммирования потока целых чисел с заданным числом
//	add := func(values []int, additive int) []int {
//		addedValues := make([]int, len(values))
//		for i, v := range values {
//			addedValues[i] = v + additive
//		}
//		return addedValues
//	}
//
//	ints := []int{1, 2, 3, 4}
//	for _, v := range multiply(multiply(add(multiply(ints, 7), 2), 1), 1) {
//		fmt.Println(v)
//	}
//}

// initStream создаёт поток целых чисел
func initStream(done <-chan int, integers ...int) <-chan int {
	output := make(chan int)
	// запуск горутины на заполнение
	go func() {
		defer close(output)
		for _, i := range integers {
			// Если передан сигнал о закрытии
			select {
			case <-done:
				return
			case output <- i:
			}
		}
	}()
	return output
}

// multiply умножает поток целых чисел на заданное число
func multiply(done <-chan int, input <-chan int, multiplier int) <-chan int {
	multipliedStream := make(chan int)
	// Запускаем горутину на умножение
	go func() {
		defer close(multipliedStream)
		// бесконечный цикл, закрывает при закрытии input (done)
		for {
			select {
			//Если передан сигнал о закрытии
			case <-done:
				return
			// Проверяем, закрыт ли канал input
			case i, isChannelOpen := <-input:
				if !isChannelOpen {
					return
				}
				// В этом случае канал открыт и выполняем операцию
				select {
				case multipliedStream <- i * multiplier:
				//	Сразу завершаем
				case <-done:
					return
				}
			}
		}

	}()
	return multipliedStream
}

// add суммирует поток целых чисел с заданным числом
func add(done <-chan int, input <-chan int, additive int) <-chan int {
	// По аналогии с умножением
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for {
			select {
			case <-done:
				return
			case i, isChannelOpen := <-input:
				if !isChannelOpen {
					return
				}
				select {
				case addedStream <- i + additive:
				case <-done:
					return
				}
			}
		}
	}()
	return addedStream
}

func main() {
	// Этот канал используется для централизованной остановки
	done := make(chan int)
	defer close(done)
	// Инициализируем поток
	intStream := initStream(done, 1, 2, 3, 4)
	// Одно из преимуществ использования каналов видно здесь -
	// каналы итерируемы, благодаря этому мы можем комбинировать
	// пайплайн как хотим, а сама целостность пайплайна от этого не меняется
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}
}
