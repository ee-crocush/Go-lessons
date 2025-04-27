package task

import (
	"github.com/ee-crocush/task-tracker/internal/domain"
	"time"
)

// TaskID идентификатор задачи.
type TaskID struct {
	value int
}

// NewTaskID создает новый идентификатор задачи.
func NewTaskID(id int) (TaskID, error) {
	if id < 1 {
		return TaskID{}, domain.ErrEmptyID
	}
	return TaskID{value: id}, nil
}

// Value возвращает значение идентификатора.
func (t TaskID) Value() int { return t.value }

// TaskTitle титул задачи.
type TaskTitle struct {
	value string
}

// NewTaskTitle создает новый титул задачи.
func NewTaskTitle(text string) (TaskTitle, error) {
	if len(text) > 0 {
		return TaskTitle{text}, nil
	}

	return TaskTitle{}, ErrInvalidLength
}

// Value возвращает значение титула.
func (t TaskTitle) Value() string { return t.value }

// TaskContent содержание задачи.
type TaskContent struct {
	value string
}

// NewTaskTitle создает содержание задачи.
func NewTaskContent(text string) (TaskContent, error) {
	if len(text) > 0 {
		return TaskContent{text}, nil
	}

	return TaskContent{}, ErrInvalidLength
}

// Value возвращает значение содержания.
func (c TaskContent) Value() string { return c.value }

// Timestamp временная метка задачи.
type Timestamp struct {
	value time.Time
}

// NewTimestamp создает новую временную метку.
func NewTimestamp() Timestamp {
	return Timestamp{time.Now().UTC()}
}

// FromUnixSeconds создаёт Timestamp из секунд
func FromUnixSeconds(s int64) Timestamp {
	return Timestamp{value: time.Unix(s, 0)}
}

func FromUnixSecondsPtr(s *int64) *Timestamp {
	if s == nil {
		return nil
	}
	return &Timestamp{value: time.Unix(*s, 0)}
}

// Time возвращает значение временной метки.
func (t Timestamp) Time() time.Time {
	return t.value
}

// String возвращает строковое значение в формате 2006-01-02 15:04:05
func (t Timestamp) String() string {
	return t.value.Format(time.DateTime)
}
