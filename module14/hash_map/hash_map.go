package main

import "fmt"

// Структура, реализующая Хеш мапу
type hashmap struct {
	data map[uint64]string // Хеш-таблица
}

// Конструктор
func newHashmap() *hashmap {
	return &hashmap{
		data: make(map[uint64]string),
	}
}

// Функция для хеширования строки
func hashstr(val string) uint64 {
	var splitted uint64

	runes := []rune(val)

	for i, r := range runes {
		splitted += uint64((i + 1) * int(r))
	}

	return splitted % 1000
}

// Метод для добавления элемента
func (h *hashmap) Set(key, val string) {
	hash := hashstr(key)
	h.data[hash] = val
}

// Метод для получения элемента
func (h *hashmap) Get(key string) (value string, ok bool) {
	hash := hashstr(key)
	value, ok = h.data[hash]
	return
}

// Метод для удаления элемента
func (h *hashmap) Delete(key string) {
	hash := hashstr(key)
	delete(h.data, hash)
}

func main() {
	hm := newHashmap()

	hm.Set("Имя", "Иван")
	hm.Set("Пол", "Мужик")
	hm.Set("Город", "Югорск")

	fmt.Println(hm.data)

	findingKey := "key1"

	val, ok := hm.Get(findingKey)
	if ok {
		fmt.Println(val)
	} else {
		fmt.Printf("Ключ '%s' не найден\n", findingKey)
	}

	hm.Delete("key5")

	hm.Set("Имя", "НеИван")
	fmt.Println(hm.data)
}
