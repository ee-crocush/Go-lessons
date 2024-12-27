// Реализация списка

package main

import (
	"fmt"
)

// Возращает размер списка
func Size(list []int) int {
	return len(list)
}

// Добавляет элемент в конец списка
func Add(list []int, elem int) []int {
	return append(list, elem)
}

// Вставляет элемент по индексу
func Insert(list []int, elem int, index int) []int {
	list = append(list, elem)

	for i := Size(list) - 1; i > index; i-- {
		list[i] = list[i-1]
	}

	list[index] = elem
	return list
}

// Удаляет элемент по индексу
func Remove(list []int, index int) []int {
	for i := index; i < Size(list)-1; i++ {
		list[i] = list[i+1]
	}
	return list[:Size(list)-1]
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(list)

	list = Add(list, 6)
	fmt.Println(list)

	list = Insert(list, 7, 3)
	fmt.Println(list)

	list = Remove(list, 3)
	fmt.Println(list)
}
