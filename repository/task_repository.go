package repository

import (
	"context"
	"hacktiv8_fp_2/entity"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task entity.Task)
	SelectTask(ctx context.Context)
	PatchTaskStatus(ctx context.Context, newStatus string)
	PatchTaskCategory(ctx context.Context, newCategory string)
	DeleteTask(ctx context.Context, id int)
}