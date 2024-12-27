// Реализуйте функцию для проверки того, что слайс отсортирован по возрастанию.
package main

import fmt "fmt"

func main() {
	initialArray := []int{1, 2, 3, 4, 5}

	fmt.Printf("Исходный массив: %v\n", initialArray)

	if checkSliceIsSorted(initialArray) {
		fmt.Println("Массив отсортирован по возрастанию")
	} else {
		fmt.Println("Массив не отсортирован по возрастанию")
	}
}

// Функция должна иметь сигнатуру:
func checkSliceIsSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
