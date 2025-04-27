package user

import "github.com/ee-crocush/task-tracker/internal/domain"

// UserID идентификатор пользователя.
type UserID struct {
	value int
}

// NewUserID создает новый идентификатор пользователя.
func NewUserID(id int) (UserID, error) {
	if id < 1 {
		return UserID{}, domain.ErrEmptyID
	}
	return UserID{value: id}, nil
}

// Value возвращает значение идентификатора пользователя.
func (u UserID) Value() int { return u.value }

// UserName имя пользователя.
type UserName struct {
	value string
}

// NewUserName создает новое имя пользователя.
// Какая-то проверка на корректность имени.
func NewUserName(name string) (UserName, error) {
	if len(name) < 6 {
		return UserName{}, ErrInvalidNameLength
	}

	return UserName{name}, nil
}

// Value возвращает значение имени пользователя.
func (n UserName) Value() string { return n.value }
