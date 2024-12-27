package main

import "fmt"

// Нахождение общих элементов двух массивов
func findCommonElements(arrayMap1 map[string]string, arrayMap2 map[string]string) []string {
	var result []string

	for i := range arrayMap1 {
		if _, ok := arrayMap2[i]; ok {
			result = append(result, i)
		}
	}
	return result
}

func printMapValues(seqName string, arrayMap map[string]string) {
	fmt.Printf("%s массив: ", seqName)
	for i := range arrayMap {
		fmt.Printf("%s ", i)
	}
	fmt.Println()
}

func main() {
	var arraySize1, arraySize2 int

	fmt.Print("Введите размер первого массива: ")
	_, err := fmt.Scan(&arraySize1) // Попытка считать целое число

	if err != nil {
		fmt.Println("Ошибка: это не число!")
		return
	}

	if arraySize1 <= 0 {
		fmt.Println("Ошибка: размер массива должен быть больше нуля!")
		return
	}

	fmt.Print("Введите размер второго массива: ")
	_, err = fmt.Scan(&arraySize2) // Попытка считать целое число

	if err != nil {
		fmt.Println("Ошибка: это не число!")
		return
	}

	if arraySize2 <= 0 {
		fmt.Println("Ошибка: размер массива должен быть больше нуля!")
		return
	}

	arrayMap1 := make(map[string]string)

	fmt.Println("Введите элементы первого массива:")

	for i := 0; i < arraySize1; i++ {
		var str string

		fmt.Scan(&str)

		arrayMap1[str] = str
	}

	arrayMap2 := make(map[string]string)

	fmt.Println("Введите элементы второго массива:")

	for i := 0; i < arraySize2; i++ {
		var str string

		fmt.Scan(&str)

		arrayMap2[str] = str
	}

	commonElements := findCommonElements(arrayMap1, arrayMap2)

	printMapValues("Первый", arrayMap1)
	printMapValues("Второй", arrayMap2)

	fmt.Println("Общие элементы:", commonElements)
}
