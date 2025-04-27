// Package user содержит доменную область пользователя.
package user

import "fmt"

// User представляет собой сущность пользователя.
type User struct {
	id   UserID
	name UserName
}

// NewUser создает новую сущность пользователя.
func NewUser(name string) (*User, error) {
	userName, err := NewUserName(name)
	if err != nil {
		return nil, fmt.Errorf("NewUser. Failed to create User: %w", err)
	}

	return &User{
		name: userName,
	}, nil
}

// ID возвращает идентификатор пользователя.
func (u *User) ID() UserID { return u.id }

// Name возвращает имя пользователя.
func (u *User) Name() UserName { return u.name }

// RehydrateUser — вспомогательный конструктор для «восстановления» сущности из БД.
func RehydrateUser(id UserID, name UserName) *User {
	return &User{
		id:   id,
		name: name,
	}
}
