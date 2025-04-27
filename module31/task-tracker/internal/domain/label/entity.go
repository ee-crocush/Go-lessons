// Package label содержит доменную область метоков задач.
package label

import "fmt"

// Label представляет собой сущность метки задачи.
type Label struct {
	id   LabelID
	name LabelName
}

// NewLabel создает новую сущность метки задачи.
func NewLabel(name string) (*Label, error) {
	labelName, err := NewLabelName(name)
	if err != nil {
		return nil, fmt.Errorf("NewLabel. Failed to create Label: %w", err)
	}

	return &Label{
		name: labelName,
	}, nil
}

// ID возвращает идентификатор метки задачи.
func (l *Label) ID() LabelID { return l.id }

// Name возвращает название метки задачи.
func (l *Label) Name() LabelName { return l.name }

// RehydrateLabel — вспомогательный конструктор для «восстановления» сущности из БД.
func RehydrateLabel(id LabelID, name LabelName) *Label {
	return &Label{
		id:   id,
		name: name,
	}
}

func (l *Label) SetID(id LabelID) { l.id = id }
