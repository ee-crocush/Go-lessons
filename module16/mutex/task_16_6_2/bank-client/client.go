package bankclient

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var _ BankClient = &Client{}

var (
	ErrNotEnoughMoney  = fmt.Errorf("на счету недостаточно средств для совершения операции")
	ErrWrongOperation  = fmt.Errorf("неверная операция. Доступные операции: balance, deposit, withdrawal, exit")
	ErrWrongBankClient = fmt.Errorf("клиент не является типом *Client")
	ErrWrongAmount     = fmt.Errorf("неверная сумма")
)

// BankClient - интерфейс клиента
type BankClient interface {
	// Deposit Пополнение депозита на указанную сумму
	Deposit(amount int)
	// Withdrawal Снятие указанной суммы со счета клиента.
	// возвращает ошибку, если баланс клиента меньше суммы снятия
	Withdrawal(amount int) error
	// Balance возвращает баланс клиента
	Balance() int
}

func NewBankClient(startDeposit int) BankClient {
	return &Client{balance: startDeposit}
}

// Client - клиент банка
type Client struct {
	mu      sync.RWMutex
	balance int
}

// Deposit - пополнение депозита
func (c *Client) Deposit(amount int) {
	c.mu.Lock()
	c.balance += amount
	c.mu.Unlock()
}

// Withdrawal - снятие денег
func (c *Client) Withdrawal(amount int) error {
	if c.balance < amount {
		return ErrNotEnoughMoney
	}

	c.mu.Lock()
	c.balance -= amount
	c.mu.Unlock()

	return nil
}

// Balance - возвращает баланс
func (c *Client) Balance() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	fmt.Printf("Текущий баланс счета: %d\n", c.balance)

	return c.balance
}

func (c *Client) ExecuteOperation() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Доступные операции: \n" +
		"1. deposit - Пополнение счета\n" +
		"2. withdrawal - Снятие средства со счета\n" +
		"3. balance - Получить текущий баланс\n" +
		"4. exit - Выход из типа банковского приложения\n")

	for {
		fmt.Print("Введите операцию: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Убираем пробелы и символы новой строки
		input = strings.ToLower(input)

		switch input {
		case "deposit":
			amount, err := readAmount(reader)

			if err != nil {
				fmt.Println(err)
				continue
			}
			c.Deposit(amount)
			fmt.Printf("Зачисление средств на счет: %d\n", amount)
		case "withdrawal":
			amount, err := readAmount(reader)

			if err != nil {
				fmt.Println(err)
				continue
			}
			err = c.Withdrawal(amount)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Снятие средств со счета: %d\n", amount)
			}
		case "balance":
			c.Balance()
		case "exit":
			fmt.Println("До свидания!")
			os.Exit(0)
		default:
			fmt.Println(ErrWrongOperation)
		}
	}
}

func readAmount(reader *bufio.Reader) (int, error) {
	fmt.Print("Введите сумму операции: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Убираем пробелы и символы новой строки
	amount, err := strconv.Atoi(input)

	if err != nil {
		return 0, ErrWrongAmount
	}

	return amount, nil
}
