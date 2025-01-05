package main

import (
	"fmt"
	"time"
)

func getChan() <-chan int {
	// Инкапсулируем инициализацию
	// и закрытие канала внутри функции
	// Далее любая попытка закрыть канал где-то
	// в другом месте (в другой горутине)
	// приведёт к ошибке уже на стадии компиляции
	c := make(chan int)
	// Горутина-замыкание
	go func() {
		defer close(c)
		for i := 1; i <= 5; i++ {
			// Отсылаем сообщения всем желающим
			// Последовательно (упорядоченная последовательность)
			// целые числа от 1 до 5
			c <- i
		}
	}()
	return c
}

//func main() {
//
//	c := getChan()
//	// Теперь принимаем сообщения
//	// и выводим их в консоль
//	for i := range c {
//		fmt.Printf("%v ", i)
//	}
//}

//const goroutineAmount int = 5
//
//func main() {
//	begin := make(chan interface{})
//	var wg sync.WaitGroup
//	// Цикл запуска пяти горутин
//	for i := 0; i < goroutineAmount; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			<-begin
//			fmt.Printf("Горутина №%d получила сигнал о закрытии и завершила свою работу\n", i)
//		}(i)
//	}
//	fmt.Println("Оповещение и разблокировка горутин произойдёт через 5 секунд...")
//	// Некий другой продолжительный по времени выполнения код,
//	// вместо которого - просто блокировка на 5 секунд
//	time.Sleep(time.Second * 5)
//	close(begin)
//	wg.Wait()
//}

//func main() {
//	// Текущее время
//	start := time.Now()
//	c := make(chan interface{})
//	go func() {
//		// Спустя пять секунд просто закрываем канал.
//		// Вспомните свойства закрытого канала
//		time.Sleep(5 * time.Second)
//		close(c)
//	}()
//	// Так как в канал ничего не отправлялось,
//	// используем в этом коде исключительно свойства закрытого
//	// канала
//	select {
//	// Соответственно, раз канал закроют примерно через 5 секунд
//	// и других case-блоков нет, - select полностью блокирует
//	// текущую горутину примерно на 5 секунд
//	case <-c:
//		fmt.Printf("Получено сообщение спустя %v\n", time.Since(start))
//	}
//}

func main() {
	var ticker *time.Ticker = time.NewTicker(time.Second * 1)
	var t time.Time
	for {
		t = <-ticker.C
		outputMessage := []byte("Время: ")
		// Метод AppendFormat преобразует объект time.Time
		// к заданному строковому формату (второй аргумент)
		// и добавляет полученную строку к строке, переданной в первом
		// аргументе
		outputMessage = t.AppendFormat(outputMessage, "15:04:05")
		fmt.Println(string(outputMessage))
	}
}
