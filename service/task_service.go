package service

import "context"

type TaskService interface {
	CreateNewTask(ctx context.Context)
	GetTasks(ctx context.Context)
	UpdateTask(ctx context.Context)
	RemoveTask(ctx context.Context)
}