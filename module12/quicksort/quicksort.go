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

func quickSortv3(a []int, low, high int) {
	if low < high {
		p := partition(a, low, high)
		quickSortv3(a, low, p-1)
		quickSortv3(a, p+1, high)
	}
}

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low

	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]

	return i
}

func quickSortv2(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar)-1
	pivotIndex := rand.Intn(len(ar))

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	quickSortv2(ar[:left])
	quickSortv2(ar[left+1:])

	return
}

func main() {
	ar := make([]int, 10)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println(ar)
	ar2 := quickSort(ar)
	fmt.Println(ar2)
}
