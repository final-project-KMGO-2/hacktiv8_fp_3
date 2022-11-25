package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"

	"github.com/mashingan/smapping"
)

type TaskService interface {
	CreateNewTask(ctx context.Context, newTask entity.TaskCreate) (entity.Task, error)
	GetTasks(ctx context.Context) ([]entity.Task, error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	UpdateTask(ctx context.Context, newTaskUpdate entity.TaskUpdate, id int) (entity.Task, error)
	ChangeTaskStatus(ctx context.Context, newStatus entity.TaskStatusModifier, id int) (entity.Task, error)
	ChangeTaskCategory(ctx context.Context, newCategory entity.TaskCategoryModifier, id int) (entity.Task, error)
	RemoveTask(ctx context.Context, id int) error
}

type taskService struct {
	taskRepo     repository.TaskRepository
	userRepo     repository.UserRepository
	categoryRepo repository.CategoryRepository
}

func NewTaskService(tr repository.TaskRepository, ur repository.UserRepository, cr repository.CategoryRepository) TaskService {
	return taskService{
		taskRepo:     tr,
		userRepo:     ur,
		categoryRepo: cr,
	}
}

func (ts taskService) CreateNewTask(ctx context.Context, newTask entity.TaskCreate) (entity.Task, error) {
	task := entity.Task{}
	err := smapping.FillStruct(&task, smapping.MapFields(&newTask))
	if err != nil {
		return entity.Task{}, err
	}
	taskCategory, err := ts.categoryRepo.GetCategoryByID(ctx, task.CategoryID)
	if err != nil || (taskCategory == entity.Category{}) {
		return entity.Task{}, errors.New("category not found")
	}
	res, err := ts.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return entity.Task{}, err
	}

	return res, nil
}
func (ts taskService) GetTasks(ctx context.Context) ([]entity.Task, error) {
	data, err := ts.taskRepo.SelectTask(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ts taskService) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	return ts.taskRepo.GetTaskByID(ctx, id)
}

func (ts taskService) UpdateTask(ctx context.Context, newTaskUpdate entity.TaskUpdate, id int) (entity.Task, error) {
	var requestMap map[string]interface{}
	data, _ := json.Marshal(newTaskUpdate)
	json.Unmarshal(data, &requestMap)
	if requestMap["category_id"] != nil {
		taskCategory, err := ts.categoryRepo.GetCategoryByID(ctx, requestMap["category_id"].(int))
		if err != nil || (taskCategory == entity.Category{}) {
			return entity.Task{}, errors.New("category not found")
		}
	}
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
	fmt.Println(requestMap)
	if requestMap["category_id"] != nil || requestMap["category_id"] != "" {
		taskCategory, err := ts.categoryRepo.GetCategoryByID(ctx, int(requestMap["category_id"].(float64)))
		if err != nil || (taskCategory == entity.Category{}) {
			return entity.Task{}, errors.New("category not found")
		}
	}
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
