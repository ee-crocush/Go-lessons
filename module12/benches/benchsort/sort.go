package benchsort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

// Реализация алгоритма сортировки пузырьком
func bubbleSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
}

// Реализация алгоритма сортировки выбором
func selectionSort(ar []int) {
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

// Реализация алгоритма сортировки вставками
func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		c := a[i]
		j := i - 1

		for j >= 0 && c < a[j] {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = c
	}
}

// Реализация алгоритма сортировки слиянием
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

// Слияние двух отсортированных массивов
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

// Реализация алгоритма быстрой сортировки
func quickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	// rand.Seed(time.Now().UnixNano())
	pivot := ar[rand.Intn(len(ar))]

	var low, equal, high []int

	for _, v := range ar {
		if v < pivot {
			low = append(low, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			high = append(high, v)
		}
	}

	sortedLow := quickSort(low)
	sortedHigh := quickSort(high)

	// Собираем результат в один массив
	result := append(sortedLow, equal...)
	result = append(result, sortedHigh...)

	return result
}
