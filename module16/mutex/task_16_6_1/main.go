package main

import (
	"fmt"
	"sync"
	"time"
)

var _ Cache = &InMemoryCache{} // это трюк для проверки типа: до тех пор пока InMemoryCache не будет реализовывать интерфейс Cache, программа не запустится

// Запись в кэше
type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

// Интерфейс Кэша
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

// Структура, которая реализует кэш в памяти
type InMemoryCache struct {
	expireIn time.Duration
	mu       sync.RWMutex
	data     map[string]*CacheEntry
}

// Конструктор создания кэша
func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn: expireIn,                     // Время жизни
		data:     make(map[string]*CacheEntry), // Данные
	}
}

// Метод для записи в кэш
func (c *InMemoryCache) Set(key string, value interface{}) {
	entry := &CacheEntry{
		settledAt: time.Now(), // Время создания
		value:     value,      // Значение
	}
	c.mu.Lock()
	c.data[key] = entry
	c.mu.Unlock()
}

// Метод для получения элемента из кэша
func (c *InMemoryCache) Get(key string) interface{} {
	// Получаем данные
	c.mu.RLock()
	entry, ok := c.data[key]
	c.mu.RUnlock()
	if !ok {
		return nil
	}

	// Проверяем интервал времени жизни
	if time.Since(entry.settledAt) > c.expireIn {
		// Удаляем запись, если время жизни истекло
		delete(c.data, key)
		return nil
	}

	return entry.value
}

func main() {
	// Создаём кэш с временем жизни записей 2 секунды
	cache := NewInMemoryCache(2 * time.Second)

	cache.Set("name", "Жора")
	cache.Set("age", 25)

	// Получаем значение из кэша
	fmt.Println("Имя:", cache.Get("name"))    // Alice
	fmt.Println("Возраст:", cache.Get("age")) // 25

	// Ждём, чтобы истёк срок действия записей
	time.Sleep(3 * time.Second)

	// Пытаемся снова получить значения
	fmt.Println("Имя после истечения срока действия:", cache.Get("name"))    // nil
	fmt.Println("Возраст после истечения срока действия:", cache.Get("age")) // nil
}