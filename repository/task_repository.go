package repository

import (
	"context"
	"hacktiv8_fp_2/entity"
	"time"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	SelectTask(ctx context.Context) ([]entity.Task, error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	UpdateTask(ctx context.Context, id int, obj map[string]interface{}) (entity.Task, error)
	DeleteTask(ctx context.Context, id int) error
}

type taskConnection struct {
	connection *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskConnection{
		connection: db,
	}
}

func (tc taskConnection) CreateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	*task.CreatedAt = time.Now()
	*task.UpdatedAt = time.Now()
	tx := tc.connection.Create(&task)
	if tx.Error != nil {
		return entity.Task{}, tx.Error
	}

	return task, nil
}

func (tc taskConnection) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var task entity.Task
	tx := tc.connection.Where(("id = ?"), id).Take(&task)
	if tx.Error != nil {
		return entity.Task{}, tx.Error
	}
	return task, nil
}

func (tc taskConnection) SelectTask(ctx context.Context) ([]entity.Task, error) {
	var tasks []entity.Task
	tx := tc.connection.Preload("User").Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tasks, nil
}

func (tc taskConnection) UpdateTask(ctx context.Context, id int, obj map[string]interface{}) (entity.Task, error) {
	var task entity.Task
	tx := tc.connection.First(&task, id)
	if tx.Error != nil {
		return entity.Task{}, tx.Error
	}
	*task.UpdatedAt = time.Now()
	tx = tc.connection.Model(&task).Updates(obj)
	if tx.Error != nil {
		return entity.Task{}, tx.Error
	}
	return task, nil
}

func (tc taskConnection) DeleteTask(ctx context.Context, id int) error {
	tx := tc.connection.Delete(&entity.Task{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
