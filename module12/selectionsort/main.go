// Реализуйте сортировку выбором, работающую «слева направо» (поиск минимальных элементов и перемещение их в начало).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 5)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	fmt.Println("Изначальный массив:")
	fmt.Println(ar)

	leftSelectionSort(ar)

	fmt.Println("Отсортированный массив:")
	fmt.Println(ar)

	ar2 := make([]int, 5)
	for i := range ar2 {
		ar2[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	fmt.Println("Изначальный массив:")
	fmt.Println(ar2)

	rightSelectionSortToDown(ar2)

	fmt.Println("Отсортированный массив от большего к меньшему:")
	fmt.Println(ar2)

	BiDirectionalSelectionSort(ar2)

	fmt.Println("Отсортированный массив от меньшего к большему:")
	fmt.Println(ar2)
}

// Реализация алгоритма сортировки выбором
func leftSelectionSort(ar []int) {
	n := len(ar)
	if n < 2 {
		return // Базовый случай: массив из 0 или 1 элемента уже отсортирован
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}

func rightSelectionSortToDown(ar []int) {
	n := len(ar)
	if n < 2 {
		return // Базовый случай: массив из 0 или 1 элемента уже отсортирован
	}
	for i := n - 1; i >= 0; i-- {
		maxIndex := i
		for j := 0; j < i; j++ {
			if ar[j] < ar[maxIndex] {
				maxIndex = j
			}
		}
		ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
	}
}

func rightSelectionSortToUp(ar []int) {
	n := len(ar)
	for i := n - 1; i >= 0; i-- {
		maxIndex := i
		for j := 0; j < i; j++ {
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}
		ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
	}
}

func BiDirectionalSelectionSort(ar []int) {
	n := len(ar)
	left, right := 0, n-1

	for left < right {
		minIndex, maxIndex := left, right
		for i := left; i <= right; i++ {
			if ar[i] < ar[minIndex] {
				minIndex = i
			}
			if ar[i] > ar[maxIndex] {
				maxIndex = i
			}

			ar[left], ar[minIndex] = ar[minIndex], ar[left]

			if maxIndex == left {
				maxIndex = minIndex
			}

			ar[right], ar[maxIndex] = ar[maxIndex], ar[right]
			left++
			right--
		}
	}
}

// Реализация алгоритма сортировки выбором через рекурсию
func selectSortRecursive(ar []int, i int) {
	n := len(ar)
	if i >= n {
		return
	}
	minIndex := i
	for j := i + 1; j < n; j++ {
		if ar[j] < ar[minIndex] {
			minIndex = j
		}
	}
	ar[i], ar[minIndex] = ar[minIndex], ar[i]
	selectSortRecursive(ar, i+1)
}
