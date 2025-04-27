package postgres

import (
	"fmt"
	"github.com/ee-crocush/task-tracker/internal/domain/task"
	"github.com/ee-crocush/task-tracker/internal/domain/user"
)

// dbTaskRow представляет строку из БД
type dbTaskRow struct {
	ID         int
	Title      string
	Content    string
	AuthorID   int
	AssignedID int
	OpenedAt   int64
	ClosedAt   *int64
}

// MapRowToTask восстанавливает сущность из бд
func MapRowToTask(row dbTaskRow) (*task.Task, error) {
	taskID, err := task.NewTaskID(row.ID)
	if err != nil {
		return nil, fmt.Errorf("mapper.MapRowToTask: %w", err)
	}

	taskTitle, err := task.NewTaskTitle(row.Title)
	if err != nil {
		return nil, fmt.Errorf("mapper.MapRowToTask: %w", err)
	}

	taskContent, err := task.NewTaskContent(row.Content)
	if err != nil {
		return nil, fmt.Errorf("mapper.MapRowToTask: %w", err)
	}

	authorUID, err := user.NewUserID(row.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("mapper.MapRowToTask: %w", err)
	}

	assignedUID, err := user.NewUserID(row.AssignedID)
	if err != nil {
		return nil, fmt.Errorf("mapper.MapRowToTask: %w", err)
	}

	oAt := task.FromUnixSeconds(row.OpenedAt)
	dAt := task.FromUnixSecondsPtr(row.ClosedAt)

	return task.RehydrateTask(taskID, taskTitle, taskContent, authorUID, assignedUID, oAt, dAt), nil
}
