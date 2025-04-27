package user

// UserRepository репозиторий, который бд должен реализовать все методы для работы с пользователями.
type UserRepository interface {
	UserCreator
	// Не будем использовать эти методы, так как они не нужны нам в рамках задачи.
	//UserFinder
	//UserUpdater
	//UserDeleter
}
