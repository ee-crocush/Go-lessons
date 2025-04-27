// Package task содержит доменную область задачи.
package task

import (
	"fmt"
	"github.com/ee-crocush/task-tracker/internal/domain/label"
	"github.com/ee-crocush/task-tracker/internal/domain/user"
)

// Task представляет собой сущность задачи.
type Task struct {
	id         TaskID
	authorID   user.UserID
	assignedID user.UserID
	title      TaskTitle
	content    TaskContent
	openedAt   Timestamp
	closedAt   *Timestamp
	labels     []*label.Label
}

// NewTask создает новую сущность задачи.
func NewTask(authorID, assignedID int, title, content string, opened int64, closed *int64) (*Task, error) {
	auID, err := user.NewUserID(authorID)
	if err != nil {
		return nil, fmt.Errorf("NewTask. Failed to create author ID: %w", err)
	}

	aID, err := user.NewUserID(assignedID)
	if err != nil {
		return nil, fmt.Errorf("NewTask. Failed to create assigned ID: %w", err)
	}

	if authorID == 0 || assignedID == 0 {
		return nil, fmt.Errorf("NewTask. Invalid user ID")
	}

	taskTitle, err := NewTaskTitle(title)
	if err != nil {
		return nil, fmt.Errorf("NewTask. Failed to create task title: %w", err)
	}

	taskContent, err := NewTaskContent(content)
	if err != nil {
		return nil, fmt.Errorf("NewTask. Failed to create task content: %w", err)
	}

	openedAt := FromUnixSeconds(opened)
	closeAt := FromUnixSecondsPtr(closed)

	return &Task{
		authorID:   auID,
		assignedID: aID,
		title:      taskTitle,
		content:    taskContent,
		openedAt:   openedAt,
		closedAt:   closeAt,
		labels:     []*label.Label{},
	}, nil
}

// ID возвращает идентификатор задачи.
func (t *Task) ID() TaskID { return t.id }

// AuthorID возвращает автора задачи.
func (t *Task) AuthorID() user.UserID { return t.authorID }

// AssignedID возвращает исполнителя задачи.
func (t *Task) AssignedID() user.UserID { return t.assignedID }

// Title возвращает заголовок задачи.
func (t *Task) Title() TaskTitle { return t.title }

// Content возвращает содержимое задачи.
func (t *Task) Content() TaskContent { return t.content }

// OpenedAt возвращает дату создания задачи.
func (t *Task) OpenedAt() Timestamp { return t.openedAt }

// ClosedAt возвращает дату закрытия задачи.
func (t *Task) ClosedAt() *Timestamp { return t.closedAt }

// Labels возвращает список меток задачи.
func (t *Task) Labels() []*label.Label { return t.labels }

// Close закрывает задачу.
func (t *Task) Close(at Timestamp) {
	t.closedAt = &at
}

// AddLabel добавляет метку к задаче.
func (t *Task) AddLabel(label *label.Label) error {
	for _, l := range t.labels {
		if l.ID() == label.ID() {
			return ErrDuplicateLabel
		}
	}

	t.labels = append(t.labels, label)

	return nil
}

// SetID устанавливает идентификатор задачи.
func (t *Task) SetID(id TaskID) { t.id = id }

// RehydrateTask — вспомогательный конструктор для «восстановления» сущности из БД.
func RehydrateTask(
	id TaskID, title TaskTitle, content TaskContent, authorID, assignedID user.UserID, openedAt Timestamp,
	closedAt *Timestamp,
) *Task {
	return &Task{
		id:         id,
		title:      title,
		content:    content,
		authorID:   authorID,
		assignedID: assignedID,
		openedAt:   openedAt,
		closedAt:   closedAt,
		labels:     []*label.Label{},
	}
}
