package repository

import (
	"context"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task entity.Task)
	SelectTask(ctx context.Context)
	PatchTaskStatus(ctx context.Context, newStatus string)
	PatchTaskCategory(ctx context.Context, newCategory string)
	DeleteTask(ctx context.Context, id int)
}

type taskConnection struct {
	connection *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskConnection{
		connection: db,
	}
}

func (tc taskConnection) CreateTask(ctx context.Context, task entity.Task)
func (tc taskConnection) SelectTask(ctx context.Context)
func (tc taskConnection) PatchTaskStatus(ctx context.Context, newStatus string)
func (tc taskConnection) PatchTaskCategory(ctx context.Context, newCategory string)
func (tc taskConnection) DeleteTask(ctx context.Context, id int)
