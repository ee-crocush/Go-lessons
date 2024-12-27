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
	ar := make([]int, 6)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println(ar)
	insertionSort(ar)

	fmt.Println(ar)
}

func insertionSortV1(a []int) {
	var n []int

	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
		n = a[:i]
	}
	a = n
	a[0] = a[len(a)-1]
}

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
