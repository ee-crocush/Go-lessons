package main

import (
	"sync"
	"time"
)

// Количество одновременно принимаемых подключений
const maxConnections int = 100

// Количество одновременных вымышленных клиентов
const amountOfClients int = 50000

// Количество обработчиков данных или размер пула потоков
const goroutinePoolSize int = 100

// Размер конвейера данных, в который попадают данные от запросов
// клиентов
const dataConveyorSize int = 100

// Data Структура представляет собой данные, которые отправляет
// клиент,
// в нашем случае это такая имитация комплексных данных
type Data int

// dataSource Буферезированный канал, в который будут направляться
// данные программе от клиентов
var dataSource chan Data

// dataConveyor Дополнительный буферизированный канал, в который
// будут направляться данные обработчикам
// Назовем его конвейер
var dataConveyor chan int

// Handler представляет собой обработчик запросов
type Handler struct {
	HandlerPool chan chan int
	DataChannel chan int
	quit        chan bool
}

// NewHandler создание нового обработчика, обратите внимание,
// HandlerPool получаем извне,
// это будет общий конвейер, как мы увидим позже, для всех
// обработчиков
func NewHandler(HandlerPool chan chan int) Handler {
	return Handler{
		HandlerPool: HandlerPool,
		DataChannel: make(chan int),
		quit:        make(chan bool)}
}

// Start метод запускает обработку данных, получаемых из общего
// конвейера
func (w Handler) Start() {
	go func() {
		for {
			// Отсылаем рабочий канал данного обработчика в общий
			// конвейер,
			// через который он будет получать данные от
			// диспетчера,
			// направляемые именно к нему на обработку
			// Отсылкой канала обработчика в общий конвейер мы как
			// бы сообщим диспетчеру,
			// что появился в наличии готовый и свободный
			// обработчик - можешь направлять очередной запрос к нему
			w.HandlerPool <- w.DataChannel
			select {
			case <-w.DataChannel:
				// имитируем обработку запроса уже без запуска в
				// отдельной горутине
				time.Sleep(time.Second)
			case <-w.quit:
				// В случае, если поступил сигнал завершения
				// работы обработчика, выходим из горутины
				return
			}
		}
	}()
}

// Stop метод предназначен для завершения работы текущего
// обработчика
func (w Handler) Stop() {
	go func() {
		w.quit <- true
	}()
}

// ============= Диспетчер конвейера обработчиков ==============

// Dispatcher - Диспетчер конвейера обработчиков
type Dispatcher struct {
	// Пул обработчиков (точнее пул индивидуальных каналов
	// обработчиков)
	HandlerPool chan chan int
}

// NewDispatcher - создание нового диспетчера
func NewDispatcher(goroutinePoolSize int) *Dispatcher {
	pool := make(chan chan int, goroutinePoolSize)
	return &Dispatcher{HandlerPool: pool}
}

// Run - Запуск определенного числа обработчиков
//
//	число обработчиков. контролируется константой
func (d *Dispatcher) Run() {
	for i := 0; i < goroutinePoolSize; i++ {
		worker := NewHandler(d.HandlerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case data := <-dataConveyor:
			// получен запрос от клиента
			go func(data int) {
				// делаем попытку получить свободный обработчик
				dataChannel := <-d.HandlerPool
				// получив внутренний канал обработчика, через
				// который он принимает данные, вверенные ему
				// отсылаем сами данные
				dataChannel <- data
			}(data)
		}
	}
}

func main() {

	dataConveyor = make(chan int, dataConveyorSize)
	// -----------------------------------------------
	startClients := func() {
		dataSource = make(chan Data, maxConnections)
		// Группа ожидания запуска клиентов
		var wgWriters sync.WaitGroup
		wgWriters.Add(amountOfClients)
		for i := 1; i <= amountOfClients; i++ {
			go func() {
				wgWriters.Done()
				dataSource <- 3
			}()
		}
		wgWriters.Wait()
	}
	// -----------------------------------------------
	// Создаем и запускаем диспетчера конвейера
	dispatcher := NewDispatcher(goroutinePoolSize)
	dispatcher.Run()
	startClients()
	for data := range dataSource {
		for i := 1; i <= int(data); i++ {
			dataConveyor <- i
		}
	}
	time.Sleep(5 * time.Minute)
}
