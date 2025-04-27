package user

import "context"

// UserCreator представляет контракт для создания пользователя.
type UserCreator interface {
	Create(ctx context.Context, user *User) (UserID, error)
}

// UserFinder представляет контракт для поиска пользователя.
// Не будет усложнять, просто ищем по id.
type UserFinder interface {
	FindByID(ctx context.Context, userID UserID) (*User, error)
}

// UserUpdater представляет контракт для обновления пользователя.
type UserUpdater interface {
	Update(ctx context.Context, user *User) error
}

// UserDeleter представляет контракт для удаления пользователя.
type UserDeleter interface {
	Delete(ctx context.Context, userID UserID) error
}
