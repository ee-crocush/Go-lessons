package task

import dom "github.com/ee-crocush/task-tracker/internal/domain/task"

// TaskDTO DTO для отображения задачи.
type TaskDTO struct {
	ID         int    `json:"id"`
	AuthorID   int    `json:"author_id"`
	AssignedID int    `json:"assigned_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	OpenedAt   int64  `json:"opened_at"`
	ClosedAt   *int64 `json:"closed_at"`
}

// LabelDTO DTO для отображения метки.
type LabelDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// toTaskDTO переводим из сущности в DTO
func toTaskDTO(t *dom.Task) TaskDTO {
	dto := TaskDTO{
		ID:         t.ID().Value(),
		AuthorID:   t.AuthorID().Value(),
		AssignedID: t.AssignedID().Value(),
		Title:      t.Title().Value(),
		Content:    t.Content().Value(),
		OpenedAt:   t.OpenedAt().Time().Unix(),
	}

	if t.ClosedAt() != nil {
		s := t.ClosedAt().Time().Unix()
		dto.ClosedAt = &s
	}

	return dto
}
