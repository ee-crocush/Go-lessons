// Реализация кольцевого массива
package main

import "fmt"

// IntRing - кольцевой массив элементов, содержащих целые числа
type IntRing struct {
	size  int    // Размер массива
	data  []*int // Содержимое массива
	Start int    // Указатель на начало массива
	End   int    // Указатель на конец массива
	// isFull bool  // Признак заполненности массива
}

// New - конструктор кольцевого массива
func New(size int, start int) (*IntRing, error) {
	if start >= size {
		return nil, fmt.Errorf("Стартовая позиция <%d> не соответствует размеру кольцевого массива <%d>", start, size)
	}

	return &IntRing{
			size:  size,
			data:  make([]*int, size, size),
			Start: start,
			End:   start,
			// isFull: false
		},
		nil
}

// Size - получение общего размера кольцевого массива
func (r IntRing) Size() int {
	return r.size
}

// IsEmpty - проверка, пуст ли кольцевой массив
func (r *IntRing) IsEmpty() bool {
	return r.Start == r.End
}

// IsFull - достигнут ли конец
func (r *IntRing) IsFull() bool {
	return (r.End < r.Start && r.End != r.Start-1) || (r.Start == 0 && r.End == r.Size()-1)
}

// Read - чтение элемента из кольцевого массива
func (r *IntRing) Read() (int, error) {
	if !r.IsEmpty() {
		el := r.data[r.Start]
		r.Start++
		for el == nil && r.Start < r.End {
			el = r.data[r.Start]
			r.Start++
		}
		if el == nil {
			return 0, fmt.Errorf("Нет новых данных в буфере")
		}

		return *el, nil
	}

	return 0, fmt.Errorf("Нет новых данных в буфере")
}

// Write - добавление одного элемента в кольцевой массив
func (r *IntRing) Write(v int) error {
	if r.IsEmpty() || !r.IsFull() {
		r.data[r.End] = &v
		r.End++

		return nil
	}

	return fmt.Errorf("Буфер полон")
}

// RemoveByIndex - удаление элемента кольцевого массива
func (r *IntRing) RemoveByIndex(index int) error {
	if index < 0 || index > r.size {
		return fmt.Errorf("<%d> - Неверный индекс удаляемого элемента кольцевого массива", index)
	}

	r.data[index] = nil

	return nil
}

func (r IntRing) Print() {
	if r.Start < r.End {
		for _, el := range r.data[r.Start:r.End] {
			if el != nil {
				fmt.Printf("%d\t", *el)
			} else {
				fmt.Printf("\t")
			}
		}
		fmt.Printf("\n")
	} else if r.Start > r.End {
		tempData := append(r.data[r.Start:], r.data[:r.End]...)
		for _, el := range tempData {
			if el != nil {
				fmt.Printf("%d\t", *el)
			} else {
				fmt.Printf("\t")
			}
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("Кольцевой массив пуст!")
	}
}

func main() {
	ring, err := New(5, 0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Размер кольцевого массива: %d\n", ring.Size())
	ring.Print()

	err = ring.Write(1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Содержимое кольцевого массива после записи элемента:")
	ring.Print()

	err = ring.Write(2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Содержимое кольцевого массива после записи элемента:")
	ring.Print()

	err = ring.Write(3)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Содержимое кольцевого массива после записи элемента:")
	ring.Print()

	err = ring.RemoveByIndex(0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Содержимое кольцевого массива после удаления элемента:")
	ring.Print()

	fmt.Println("Читаем каждый элемент кольцевого массива:")

	for i := 0; i < ring.Size(); i++ {
		el, err := ring.Read()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%d (Указатель: %d)\n", el, ring.Start)
	}

	err = ring.Write(4)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Содержимое кольцевого массива после записи элемента:")
	ring.Print()
	fmt.Printf(
		"Размер кольцевого массива: %d, Указатель на начало: %d, Указатель на конец: %d\n",
		ring.Size(),
		ring.Start,
		ring.End,
	)

	err = ring.Write(5)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Содержимое кольцевого массива после записи элемента:")
	ring.Print()
	fmt.Printf(
		"Размер кольцевого массива: %d, Указатель на начало: %d, Указатель на конец: %d\n",
		ring.Size(),
		ring.Start,
		ring.End,
	)

	err = ring.Write(6)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Содержимое кольцевого массива после записи элемента:")
	ring.Print()
	fmt.Printf(
		"Размер кольцевого массива: %d, Указатель на начало: %d, Указатель на конец: %d\n",
		ring.Size(),
		ring.Start,
		ring.End,
	)
}
