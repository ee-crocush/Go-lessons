package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

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

func checkSliceIsSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func main() {
	tests := [][]int{
		{0, 1, 2, 3, 4, 5},
		{9, 7, 4, 1, 3, 5},
		{0},
		{},
		{1, 1},
		{3, 2, 1},
		{5, 15, 2, 13, 7, 16, 10, 2},
		{1, 9, 7, 4, 6, 2, 1, 13, 22, -3, 12, 76},
	}
	for _, test := range tests {
		result := quickSort(test)
		if !checkSliceIsSorted(result) {
			fmt.Errorf("Массив %v не отсортирован по возрастанию", result)
		} else {
			fmt.Printf("Массив %v отсортирован по возрастанию, все заебись\n", result)
		}
	}
}
