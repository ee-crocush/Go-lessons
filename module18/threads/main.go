package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Количество одновременных запросов от вымышленных клиентов
const amountOfClients int = 500000

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
	printMemUsage("Начальное состояние памяти:")
	// Вспомогательный канал, который нам потребуется для
	// синхронизации горутины,
	// которая будет делать ориентировочные замеры времени,
	// затраченного на принятие всех запросов
	timeMeasuringRoutineStarted := make(chan bool)
	var wgClientsDone sync.WaitGroup
	wgClientsDone.Add(amountOfClients)
	getDataSource := func() chan int {
		dataSource := make(chan int, 100)
		// Вспомогательная группа ожидания,
		// которая нужна этому замыканию,
		// чтобы убедиться что все клиенты
		// созданы и готовы к работе перед выходом из этого замыкания
		var wgClients sync.WaitGroup
		wgClients.Add(amountOfClients)
		for i := 1; i <= amountOfClients; i++ {
			go func() {
				defer wgClientsDone.Done()
				wgClients.Done()
				dataSource <- 2
			}()
		}
		// Ждем момента, когда все клиенты запустятся
		wgClients.Wait()
		return dataSource
	}
	dataSourceChan := getDataSource()
	printMemUsage("После создания клиентов:")
	// Горутина, которая дождется завершения работы клиентов и
	// закроет
	// канал запросов. В ней же происходит замер времени, которое
	// требуется на
	// получение всех запросов (обратите внимание не на фактическую
	// обработку,
	// а на фактическое получение от клиентов всех запросов)
	go func() {
		start := time.Now()
		timeMeasuringRoutineStarted <- true
		wgClientsDone.Wait()
		fmt.Printf("Время обработки запросов: %d мс\n", time.Since(start).Milliseconds())
		close(dataSourceChan)
	}()
	// Запускаем горутину для замера времени
	// и других вспомогательных операций (закрытия основного канала
	// запросов)
	<-timeMeasuringRoutineStarted
	for data := range dataSourceChan {
		// Запускаем экземпляр обработчика данных
		go func(data int) {
			// Обработка самого запроса
			for i := 1; i <= data; i++ {
				go func() {
					var i int = 0
					// Имитация работы обработчика запроса
					for i <= 50000 {
						i++
					}
				}()
			}
		}(data)
	}
	printMemUsage("После обработки запросов:")
	time.Sleep(5 * time.Second)
	printMemUsage("Конечное состояние памяти:")
}
