package author

// Repository репозиторий представляющий контракты для работы с авторами.
type Repository interface {
	Creator
	Finder
	Writer
}
