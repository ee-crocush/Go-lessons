package utils

import "strings"

// ExtractErrorMessage извлекает последнее сообщение из цепочки ошибок.
func ExtractErrorMessage(err error) string {
	parts := strings.Split(err.Error(), ": ")
	return parts[len(parts)-1] // Берем последнее
}
