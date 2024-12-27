package bankclient

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var _ BankClient = &Client{}

var (
	ErrNotEnoughMoney  = fmt.Errorf("на счету недостаточно средств для совершения операции")
	ErrWrongOperation  = fmt.Errorf("неверная операция. Доступные операции: balance, deposit, withdrawal, exit")
	ErrWrongBankClient = fmt.Errorf("клиент не является типом *Client")
	ErrWrongAmount     = fmt.Errorf("неверная сумма")
	ErrWrongRange      = fmt.Errorf("неверный диапазон минимального и максимального значения суммы")
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

// BankClient - интерфейс клиента
type BankClient interface {
	// Пополнение депозита на указанную сумму
	Deposit(amount int)
	// Снятие указанной суммы со счета клиента.
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
	c.mu.RLock()
	c.balance += amount
	c.mu.RUnlock()
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

func (c *Client) ExecuteOperation(minVal, maxVal int) {
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
			amount := randomAmount(minVal, maxVal)
			c.Deposit(amount)
			fmt.Printf("Зачисление средств на счет: %d\n", amount)
		case "withdrawal":
			amount := randomAmount(minVal, maxVal)
			c.Withdrawal(amount)
			fmt.Printf("Снятие средств со счета: %d\n", amount)
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

// randomAmount - генерирует случайную сумму средств
func randomAmount(minVal, maxVal int) int {
	if minVal >= maxVal {
		fmt.Println(ErrWrongRange)

		return 0
	}

	amount := rand.Intn(maxVal-minVal+1) + minVal

	return amount
}
