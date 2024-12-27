package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}
	fmt.Println("Исходный массив:")
	fmt.Println(ar)

	// spaws := bubbleSort(ar)
	swaps := 0
	bubbleSortRecursive(ar, &swaps)
	fmt.Println("Отсортированный массив:")
	fmt.Println(ar)
	fmt.Printf("Количество перестановок: %d\n", swaps)

	if swaps == 0 {
		fmt.Println("Массив был изначально отсортирован!")
		os.Exit(0)
	}

	bubbleSortReverse(ar)
	fmt.Println("Отсортированный массив в обратном порядке:")
	fmt.Println(ar)

}

func bubbleSort(ar []int) int {
	var spaws int

	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if ar[i] > ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
				spaws++
			}
		}
	}

	return spaws
}

func bubbleSortRecursive(ar []int, swaps *int) {
	n := len(ar)
	if n < 2 {
		return // Базовый случай: массив из 0 или 1 элемента уже отсортирован
	}

	for i := 0; i < n-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i], ar[i+1] = ar[i+1], ar[i]
			*swaps++
		}
	}

	bubbleSortRecursive(ar[:n-1], swaps)
}

func bubbleSortReverse(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if ar[i] < ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
}
