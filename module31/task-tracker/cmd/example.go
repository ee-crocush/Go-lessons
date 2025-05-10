package main

import (
	"context"
	"fmt"
	"github.com/ee-crocush/task-tracker/internal/infrastructure/repository/postgres"
	uclabel "github.com/ee-crocush/task-tracker/internal/usecase/label"
	uctask "github.com/ee-crocush/task-tracker/internal/usecase/task"
	ucuser "github.com/ee-crocush/task-tracker/internal/usecase/user"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

// CreateLabel - создание новой метки в БД.
func CreateLabel(ctx context.Context, pool *pgxpool.Pool) (uclabel.CreateOutputDTO, error) {
	repo := postgres.NewLabelPGRepository(pool)

	uc := uclabel.NewCreateUseCase(repo)

	// Генерируем случайный uuid
	randomUUID := uuid.New()

	in := uclabel.CreateInputDTO{
		Name: fmt.Sprintf("Метка-%s", randomUUID.String()),
	}
	out, err := uc.Execute(ctx, in)
	if err != nil {
		return uclabel.CreateOutputDTO{}, err
	}

	return out, nil
}

// CreateUser - создание нового пользователя в БД.
func CreateUser(ctx context.Context, pool *pgxpool.Pool) (int, error) {
	repo := postgres.NewUserPGRepository(pool)

	uc := ucuser.NewCreateUseCase(repo)

	in := ucuser.CreateInputDTO{
		Name: "Иван Иванов",
	}
	out, err := uc.Execute(ctx, in)
	if err != nil {
		return 0, err
	}
	return out.ID, nil
}

// CreateTask - создание новой задачи.
// Передаем ID автора и исполнителя и массив меток.
func CreateTask(
	ctx context.Context, pool *pgxpool.Pool, authorID, assignedID int,
	labels []uctask.LabelDTO,
) (int, error) {
	repo := postgres.NewTaskPGRepository(pool)

	uc := uctask.NewCreateUseCase(repo)

	// Генерируем случайный uuid
	randomUUID := uuid.New()

	title := fmt.Sprintf("Задача для примера-%s", randomUUID.String())
	content := fmt.Sprintf("Описание задачи для примера-%s.", randomUUID.String())

	in := uctask.CreateInputDTO{
		AuthorID:   authorID,
		AssignedID: assignedID,
		Title:      title,
		Content:    content,
		Labels:     labels,
	}
	out, err := uc.Execute(ctx, in)
	if err != nil {
		return 0, err
	}
	return out.ID, nil
}

// GetTaskByID Получает задачу по ее ID.
func GetTaskByID(ctx context.Context, pool *pgxpool.Pool, id int) (
	uctask.TaskDTO, error,
) {
	repo := postgres.NewTaskPGRepository(pool)

	uc := uctask.NewFindByIDUseCase(repo)

	in := uctask.FindByIDInputDTO{
		ID: id,
	}
	out, err := uc.Execute(ctx, in)
	if err != nil {
		return uctask.TaskDTO{}, err
	}
	return out, nil
}

// GetTasksByAuthor - получение списка задач по автору.
// Передаем ID автора
func GetTasksByAuthor(ctx context.Context, pool *pgxpool.Pool, authorID int) (
	uctask.FindAllByAuthorIDOutputDTO, error,
) {
	repo := postgres.NewTaskPGRepository(pool)

	uc := uctask.NewFindAllByAuthorIDUseCase(repo)

	in := uctask.FindAllByAuthorIDDTO{
		AuthorID: authorID,
	}
	out, err := uc.Execute(ctx, in)
	if err != nil {
		return uctask.FindAllByAuthorIDOutputDTO{}, err
	}
	return out, nil
}

// GetAllTasks получает все задачи.
func GetAllTasks(ctx context.Context, pool *pgxpool.Pool) (
	uctask.FindAllOutputDTO, error,
) {
	repo := postgres.NewTaskPGRepository(pool)

	uc := uctask.NewFindAllUseCase(repo)

	out, err := uc.Execute(ctx)
	if err != nil {
		return uctask.FindAllOutputDTO{}, err
	}
	return out, nil
}

// GetTasksByLabel - получение списка задач по метке.
// Передаем ID метки
//func GetTasksByLabel(ctx context.Context, pool *pgxpool.Pool, labelID int) (
//	uctask.Fin, error,
//) {
//	repo := postgres.NewTaskPGRepository(pool)
//
//	uc := uctask.NewFindAllByAuthorIDUseCase(repo)
//
//	in := uctask.FindAllByAuthorIDDTO{
//		AuthorID: authorID,
//	}
//	out, utils := uc.Execute(ctx, in)
//	if utils != nil {
//		return uctask.FindAllByAuthorIDOutputDTO{}, utils
//	}
//	return out, nil
//}
