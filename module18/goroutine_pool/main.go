package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// WorkerPool представляет пул потоков.
type WorkerPool struct {
	tasks      chan func()    // Канал для задач
	workerDone sync.WaitGroup // Группа ожидания для всех воркеров
}

// NewWorkerPool создает новый пул потоков.
func NewWorkerPool(poolSize int) *WorkerPool {
	return &WorkerPool{
		tasks: make(chan func(), poolSize), // Буферизированный канал задач
	}
}

// Start запускает пул воркеров.
func (wp *WorkerPool) Start(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		wp.workerDone.Add(1)
		go wp.worker(i)
	}
}

// worker выполняет задачи из канала до его закрытия.
func (wp *WorkerPool) worker(workerID int) {
	defer wp.workerDone.Done()
	for task := range wp.tasks {
		//fmt.Printf("Worker %d processing task\n", workerID)
		task()
	}
	fmt.Printf("Worker %d stopped\n", workerID)
}

// Submit добавляет задачу в пул для выполнения.
func (wp *WorkerPool) Submit(task func()) {
	wp.tasks <- task
}

// Stop завершает выполнение пула, ожидая выполнения всех задач.
func (wp *WorkerPool) Stop() {
	close(wp.tasks)      // Закрываем канал задач
	wp.workerDone.Wait() // Ждем завершения всех воркеров
	fmt.Println("All workers have stopped")
}

// printMemUsage выводит информацию о текущем состоянии памяти.
func printMemUsage(message string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s\n", message)
	fmt.Printf("Alloc = %v KiB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc = %v KiB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys = %v KiB\n", m.Sys/1024)
	fmt.Printf("NumGC = %v\n\n", m.NumGC)
}

func main() {
	// Создаем пул воркеров с размером канала задач
	poolSize := 10
	numWorkers := 100

	printMemUsage("Начальное состояние памяти:")
	wp := NewWorkerPool(poolSize)
	wp.Start(numWorkers)
	printMemUsage("После запуска пула:")
	// Отправляем задачи в пул
	numTasks := 100000
	for i := 0; i < numTasks; i++ {
		//taskID := i
		wp.Submit(func() {
			time.Sleep(100 * time.Millisecond) // Имитация работы
			//fmt.Printf("Task %d is complete\n", taskID)
		})
	}

	printMemUsage("После завершения работы пула:")
	// Завершаем работу пула
	wp.Stop()
}
