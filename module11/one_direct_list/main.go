// Реализация односвязного списка

package main

import (
	"fmt"
)

var ErrorWrongListIndex = fmt.Errorf("Неверный индекс списка")

// IntNode - описание типа Узел списка
type IntNode struct {
	Value int
	Next  *IntNode
}

// New - создание нового узла списка
func New(value int) *IntNode {
	return &IntNode{value, nil}
}

// IntList - описание типа Список целых чисел
type IntList struct {
	size int
	Head *IntNode
}

// Size - получение размера списка
func (l IntList) Size() int {
	return l.size
}

// Get - получение узла по индексу
func (l IntList) Get(index int) (*IntNode, error) {
	if index < 0 || index >= l.Size() {
		return nil, ErrorWrongListIndex
	}

	node := l.Head

	for i := 1; i <= index; i++ {
		node = node.Next
	}

	return node, nil
}

// Set - обновление произвольного элемента списка
func (l *IntList) Set(el int, index int) error {
	if index < 0 || index >= l.Size() {
		return ErrorWrongListIndex
	}

	node, err := l.Get(index)

	if err != nil {
		return err
	}

	node.Value = el

	return nil
}

// Add - добавление нового элемента в начало списка
func (l *IntList) Add(el int) {
	newNode := New(el)

	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}

	l.size++

}

// Insert - вставка нового элемента в произвольную позицию
func (l *IntList) Insert(el int, index int) error {
	if index < 0 || index >= l.Size() {
		return ErrorWrongListIndex
	}

	newNode := New(el)

	if index == 0 {
		l.Add(el)

		return nil
	}

	node, err := l.Get(index - 1)

	if err != nil {
		return err
	}

	newNode.Next = node.Next
	node.Next = newNode
	l.size++

	return nil
}

// Remove - удаление элемента из проивольной позиции
func (l *IntList) Remove(index int) error {
	if index < 0 || index >= l.Size() {
		return ErrorWrongListIndex
	}

	if index == 0 {
		l.Head = l.Head.Next
	} else {
		node, err := l.Get(index - 1)

		if err != nil {
			return err
		}

		node.Next = node.Next.Next
	}

	l.size--

	return nil
}

// Print - печать списка
func (l IntList) Print() {
	node := l.Head

	if node != nil {
		for node != nil {
			fmt.Printf("%d\t", node.Value)
			node = node.Next
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("Список пуст!")
	}
}

func main() {
	list := IntList{}
	list.Print()
	list.Add(2)
	list.Add(1)
	list.Add(0)
	list.Print()
	list.Insert(8, 1)
	list.Print()
	list.Remove(0)
	list.Print()
	fmt.Println(list.Size())
}
