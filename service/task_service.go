package service

import (
	"context"
	"hacktiv8_fp_2/repository"
)

type TaskService interface {
	CreateNewTask(ctx context.Context)
	GetTasks(ctx context.Context)
	UpdateTask(ctx context.Context)
	RemoveTask(ctx context.Context)
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskRepository(tr repository.TaskRepository) TaskService{
	return taskService{
		taskRepo: tr,
	}
}

func (ts taskService) CreateNewTask(ctx context.Context)
func (ts taskService) GetTasks(ctx context.Context)
func (ts taskService) UpdateTask(ctx context.Context)
func (ts taskService) RemoveTask(ctx context.Context)