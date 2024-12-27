package main

import (
	"fmt"
	"math/rand"
	"sync"
	bankclient "task_16_6_2/bank-client"
	"time"
)

var ErrWrongRange = fmt.Errorf("неверный диапазон минимального и максимального значения суммы")

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func randomAmount(minVal, maxVal int) int {
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

			realClient.ExecuteOperation(minDeposit, maxDeposit)
		}()
	} else {
		fmt.Println(bankclient.ErrWrongBankClient)
	}

	// Запускаем 10 горутин для пополнения счета
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			amount := randomAmount(minDeposit, maxDeposit)

			client.Deposit(amount)

			// Задержка между пополнением от 0.5 с до 1 с
			time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		}()
	}

	// Запускаем 5 горутин для уменьшения баланса
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			amount := randomAmount(minWithdrawal, maxWithdrawal)

			err := client.Withdrawal(amount)

			if err != nil {
				fmt.Println(err)
			}

			// Задержка между пополнением от 0.5 с до 1 с
			time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		}()
	}

	// for {

	wg.Wait()
	// }

}
