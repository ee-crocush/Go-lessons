package semaphore

import (
	"fmt"
	"time"
)

// Semaphore — структура двоичного семафора
type Semaphore struct {
	sem     chan struct{} // Канал для управления доступом
	timeout time.Duration // Таймаут для операций с семафором
}

// NewSemaphore — функция создания семафора
func NewSemaphore(timeout time.Duration) *Semaphore {
	return &Semaphore{
		sem:     make(chan struct{}, 1), // Буфер 1 для двоичного семафора
		timeout: timeout,
	}
}

// Acquire — метод захвата семафора
func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- struct{}{}: // Успешно захватили
		return nil
	case <-time.After(s.timeout): // Таймаут истёк
		return fmt.Errorf("Не удалось захватить семафор: таймаут")
	}
}

// Release — метод освобождения семафора
func (s *Semaphore) Release() error {
	select {
	case <-s.sem: // Успешно освободили
		return nil
	case <-time.After(s.timeout): // Таймаут истёк
		return fmt.Errorf("Не удалось освободить семафор: таймаут")
	}
}
