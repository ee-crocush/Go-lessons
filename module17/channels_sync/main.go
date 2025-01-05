package main

import (
	"fmt"
	"sync"
)

// Количество транзакций, которые надо совершить над банковским
// счетом
const amountOfTransactions int = 1000

// Функция выполнения транзакции
func doTransaction(summ int, bank chan<- int, transactionDone <-chan int, wg *sync.WaitGroup) int {
	defer wg.Done()
	bank <- summ
	return <-transactionDone
}

// Функция создания банка
func createBank(bankAccount *int, wg *sync.WaitGroup) (bank chan int, transactionDone chan int) {
	bank = make(chan int)
	transactionDone = make(chan int)
	go func() {
		defer wg.Done()
		for summ := range bank {
			*bankAccount = *bankAccount + summ
			transactionDone <- *bankAccount
		}
	}()
	return
}

func main() {
	// Счет банка
	var bankAccount int = 0
	// Количество горутин, необходимых для выполнения всех
	// транзакций
	var amountOfGoRoutines int = amountOfTransactions * 2
	var wg sync.WaitGroup
	wg.Add(amountOfGoRoutines)
	bank, transactionDone := createBank(&bankAccount, &wg)
	for i := 1; i <= amountOfTransactions; i++ {
		//fmt.Printf("Транзакция #%d\n", i)
		go doTransaction(2, bank, transactionDone, &wg)
		go doTransaction(-1, bank, transactionDone, &wg)
	}
	// Дожидаемся завершения всех транзакций
	wg.Wait()
	wg.Add(1)
	close(bank)
	// Дожидаемся закрытия банка
	wg.Wait()
	fmt.Println(bankAccount)
}
