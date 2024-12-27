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
	ar := make([]int, 10)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println(ar)
	ar = mergeSort(ar)
	fmt.Println(ar)
}

func mergeSort(arr []int) []int {
	// Условие выхода: если массив состоит из одного или нуля элементов, он уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Разделение массива на две половины
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Слияние отсортированных половин
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Сравнение первых элементов каждого массива и добавление меньшего в результат
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавление оставшихся элементов из left и right, если они остались
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
