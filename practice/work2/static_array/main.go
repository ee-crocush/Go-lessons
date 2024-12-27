package main

import (
	"fmt"
	"strconv"
)

func Print(t []int) {
	fmt.Println("Температуры на станциях: ")

	for i := 0; i < len(t); i++ {
		fmt.Printf("%d\t", t[i])
	}
	fmt.Print("\n")
}

func Average(t []int) float32 {
	var total int

	for _, value := range t {
		total += value
	}
	return float32(total) / float32(len(t))
}

func getValidInput(prompt string, min int, max int) (int, error) {
	for {
		var input string

		fmt.Print(prompt)
		fmt.Scanln(&input)

		if input == "exit" {
			return 0, fmt.Errorf("Выход из программы")
		}

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите корректное целое число.")
			continue
		}

		// Проверка на диапазон, если требуется
		if max > 0 && (value < min || value > max) {
			fmt.Printf("Ошибка: введите значение от %d до %d.\n", min, max)
			continue
		}
		return value, nil
	}
}

func InputTemperature(temperatures []int) {
	fmt.Println("Напишите 'exit' для выхода из программы")
	fmt.Printf("Всего станций: %d\n", len(temperatures))

	for {
		station, err := getValidInput("Введите номер метеостанции (начиная с 1): ", 1, len(temperatures))
		if err != nil {
			fmt.Println(err)
			break
		}

		temperature, err := getValidInput("Введите температуру: ", 0, 0)
		if err != nil {
			fmt.Println(err)
			break
		}

		temperatures[station-1] = temperature
		average := Average(temperatures)

		fmt.Printf("Температура %d-й станции обновлена до: %d\n", station, temperature)
		fmt.Printf("Средняя температура: %.2f\n", average)
	}
}

func main() {
	meteo := make([]int, 5, 5)

	for i := 0; i < len(meteo); i++ {
		fmt.Printf("Введите температуры %d-й станции: ", i+1)
		fmt.Scanln(&meteo[i])
	}

	Print(meteo)
	InputTemperature(meteo)
}
