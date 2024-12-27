package main

import testing "testing"

func TestSort(t *testing.T) {
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
			t.Errorf("Массив %v не отсортирован по возрастанию", result)
		} else {
			t.Logf("Массив %v отсортирован по возрастанию, все заебись", result)
		}
	}
}

// Функция провяет, что массив отсотирован по возрастанию
// func checkSliceIsSorted(a []int) bool {
// 	for i := 0; i < len(a)-1; i++ {
// 		if a[i] > a[i+1] {
// 			return false
// 		}
// 	}
// 	return true
// }
