package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 5, 2, 2 - 1, -2, -3}
	min, err := findMaxNegative(array)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(min)

	mostOften, err := findMostOftenRepeated(array)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mostOften)

	mostOftenWithMap, err := findMostOftenRepeatedWithMap(array)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mostOftenWithMap)

	nonegative := trimNegative(array)
	fmt.Println(nonegative)

	var array2 = []int{}

	moreAverage, err := moreThanAverage(array2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(moreAverage)
}

func findMaxNegative(array []int) (min int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found min in empty slice")
	}

	min = array[0]
	for _, val := range array[1:] {
		if val < min {
			min = val
		}
	}

	if min > 0 {
		return min, fmt.Errorf("could not found min in negative slice, only min value in slice")
	}

	return min, nil
}

func findMostOftenRepeated(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for j := 1; j < len(array); j++ {
			if number == array[j] {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}

func findMostOftenRepeatedWithMap(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var countMap = make(map[int]int)

	for _, number := range array {
		countMap[number]++
	}

	var maxIndex, maxCount = 0, 0

	for i, number := range array {
		if countMap[number] > maxCount {
			maxIndex = i
			maxCount = countMap[number]
		}
	}

	return array[maxIndex], nil
}

func trimNegative(array []int) []int {
	result := make([]int, 0, len(array))

	for _, val := range array {
		if val >= 0 {
			result = append(result, val)
		}
	}

	return result
}

func moreThanAverage(array []int) (result []int, err error) {
	if len(array) == 0 {
		return result, fmt.Errorf("could not found average in empty slice")
	}

	var sum float32

	for _, val := range array {
		sum += float32(val)
	}

	average := sum / float32(len(array))

	for _, val := range array {
		if val > int(average) {
			result = append(result, val)
		}
	}

	return result, nil
}
