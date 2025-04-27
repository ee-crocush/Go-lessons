package label

import (
	"github.com/ee-crocush/task-tracker/internal/domain"
)

// LabelID идентификатор метки.
type LabelID struct {
	value int
}

// NewLabelID создает новый идентификатор метки
func NewLabelID(id int) (LabelID, error) {
	if id < 1 {
		return LabelID{}, domain.ErrEmptyID
	}
	return LabelID{value: id}, nil
}

// Value возвращает значение идентификатора метки.
func (l LabelID) Value() int { return l.value }

// LabelName представляет собой наименование метки.
type LabelName struct {
	name string
}

// NewLabelName создает новое наименование метки.
func NewLabelName(name string) (LabelName, error) {
	if len(name) > 0 {
		return LabelName{name}, nil
	}

	return LabelName{}, ErrInvalidEmptyLabel
}

// Value возвращает значение наименования метки.
func (n LabelName) Value() string { return n.name }
