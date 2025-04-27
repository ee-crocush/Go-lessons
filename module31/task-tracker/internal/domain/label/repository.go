package label

// LabelRepository репозиторий, который бд должен реализовать все методы для работы с меткими.
// Не будем усложнять интерфейс, будем реализовывать только создание
type LabelRepository interface {
	LabelCreator
}
