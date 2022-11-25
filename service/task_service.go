package service

import (
	"context"
	"encoding/json"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"
	"time"

	"github.com/mashingan/smapping"
)

type TaskService interface {
	CreateNewTask(ctx context.Context, newTask entity.TaskCreate) (entity.Task, error)
	GetTasks(ctx context.Context) ([]entity.TaskDetail, error)
	UpdateTask(ctx context.Context, newTaskUpdate entity.TaskUpdate, id int) (entity.Task, error)
	ChangeTaskStatus(ctx context.Context, newStatus entity.TaskStatusModifier, id int) (entity.Task, error)
	ChangeTaskCategory(ctx context.Context, newCategory entity.TaskCategoryModifier, id int) (entity.Task, error)
	RemoveTask(ctx context.Context, id int) error
}

type taskService struct {
	taskRepo repository.TaskRepository
	userRepo repository.UserRepository
}

func NewTaskService(tr repository.TaskRepository, ur repository.UserRepository) TaskService {
	return taskService{
		taskRepo: tr,
		userRepo: ur,
	}
}

func (ts taskService) CreateNewTask(ctx context.Context, newTask entity.TaskCreate) (entity.Task, error) {
	task := entity.Task{}
	err := smapping.FillStruct(&task, smapping.MapFields(&newTask))
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	if err != nil {
		return entity.Task{}, err
	}

	res, err := ts.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return entity.Task{}, err
	}

	return res, nil
}
func (ts taskService) GetTasks(ctx context.Context) ([]entity.TaskDetail, error) {
	data, err := ts.taskRepo.SelectTask(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (ts taskService) UpdateTask(ctx context.Context, newTaskUpdate entity.TaskUpdate, id int) (entity.Task, error) {
	var requestMap map[string]interface{}
	data, _ := json.Marshal(newTaskUpdate)
	json.Unmarshal(data, &requestMap)

	res, err := ts.taskRepo.UpdateTask(ctx, id, requestMap)

	if err != nil {
		return entity.Task{}, err
	}
	return res, nil
	// ubah struct menjadi map string lalu pass ke repo
}
func (ts taskService) ChangeTaskStatus(ctx context.Context, newStatus entity.TaskStatusModifier, id int) (entity.Task, error) {
	var requestMap map[string]interface{}
	data, _ := json.Marshal(newStatus)
	json.Unmarshal(data, &requestMap)

	res, err := ts.taskRepo.UpdateTask(ctx, id, requestMap)

	if err != nil {
		return entity.Task{}, err
	}
	return res, nil
}
func (ts taskService) ChangeTaskCategory(ctx context.Context, newCategory entity.TaskCategoryModifier, id int) (entity.Task, error) {
	var requestMap map[string]interface{}
	data, _ := json.Marshal(newCategory)
	json.Unmarshal(data, &requestMap)

	res, err := ts.taskRepo.UpdateTask(ctx, id, requestMap)

	if err != nil {
		return entity.Task{}, err
	}
	return res, nil
}

func (ts taskService) RemoveTask(ctx context.Context, id int) error {
	err := ts.taskRepo.DeleteTask(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
