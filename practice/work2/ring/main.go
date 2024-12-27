package main

import (
	"container/ring"
	"fmt"
)

// Размер кольцевого буфера
const bufferSize = 10

// encode симулирует "кодирование" данных перед отправкой
func encode(data string) string {
	return data // В данной задаче просто возвращаем исходные данные
}

// decode симулирует "декодирование" данных после получения
func decode(data string) string {
	return data // Аналогично, просто возвращаем исходные данные
}

// simulateTransmission имитирует передачу данных через кольцевой буфер
func simulateTransmission(input string) string {
	// Инициализируем кольцевой буфер нужного размера
	buffer := ring.New(bufferSize)

	// Разбиваем строку на части, размером равным bufferSize (для примера)
	parts := len(input) / bufferSize
	if len(input)%bufferSize != 0 {
		parts++ // Увеличиваем на 1, если есть остаток
	}

	// Отправляем данные по частям
	for i := 0; i < parts; i++ {
		start := i * bufferSize
		end := start + bufferSize
		if end > len(input) {
			end = len(input)
		}

		// "Кодируем" и записываем данные в буфер
		encodedData := encode(input[start:end])
		buffer.Value = encodedData
		buffer = buffer.Next()
	}

	// Принимаем данные из буфера и "декодируем"
	var output string
	buffer.Do(func(value interface{}) {
		if value != nil {
			decodedData := decode(value.(string))
			output += decodedData
		}
	})

	return output
}

func main() {
	var input string

	fmt.Println("Введите сообщение (для выхода напишите 'exit'):")

	for {
		fmt.Print("> ")
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("Завершение программы.")
			break
		}

		result := simulateTransmission(input)
		fmt.Println("Получено сообщение:", result)
	}
}
