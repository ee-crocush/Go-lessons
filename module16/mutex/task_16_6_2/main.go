package main

import (
	"fmt"
	"math/rand"
	"sync"
	bankclient "task_16_6_2/bank-client"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomAmount(minVal, maxVal int) int {
	var ErrWrongRange = fmt.Errorf("неверный диапазон минимального и максимального значения суммы")
	if minVal >= maxVal {
		fmt.Println(ErrWrongRange)
		return 0
	}

	amount := rand.Intn(maxVal-minVal+1) + minVal

	return amount
}

func main() {
	// Создаем клиента через интерфейс BankClient
	client := bankclient.NewBankClient(100)

	// Указываем диапазон сумм
	minDeposit := 1
	maxDeposit := 10
	minWithdrawal := 1
	maxWithdrawal := 5

	var wg sync.WaitGroup

	if realClient, ok := client.(*bankclient.Client); ok {
		wg.Add(1)

		go func() {
			defer wg.Done()
			realClient.ExecuteOperation()
		}()
	} else {
		fmt.Println(bankclient.ErrWrongBankClient)
	}

	// Запускаем 10 горутин для пополнения счета
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			amount := RandomAmount(minDeposit, maxDeposit)
			// Задержка между пополнением от 0.5 с до 1 с
			time.Sleep(time.Duration(i*rand.Intn(500)+500) * time.Millisecond)
			client.Deposit(amount)
		}(i)
	}

	// Запускаем 5 горутин для уменьшения баланса
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			amount := RandomAmount(minWithdrawal, maxWithdrawal)
			// Задержка между пополнением от 0.5 с до 1 с
			time.Sleep(time.Duration(5*rand.Intn(500)+500) * time.Millisecond)

			err := client.Withdrawal(amount)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
}
