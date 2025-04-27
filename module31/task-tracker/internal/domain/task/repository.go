package task

// TaskRepository репозиторий, который бд должен реализовать все методы для работы с задачами.
type TaskRepository interface {
	TaskCreator
	TaskFinder
	TaskUpdater
	TaskDeleter
}
