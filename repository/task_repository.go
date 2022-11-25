package repository

import (
	"context"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	SelectTask(ctx context.Context) ([]entity.TaskDetail, error)
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
	tx := tc.connection.Create(&task)
	if tx.Error != nil {
		return entity.Task{}, tx.Error
	}

	return task, nil
}
func (tc taskConnection) SelectTask(ctx context.Context) ([]entity.TaskDetail, error) {
	var tasks []entity.TaskDetail
	tx := tc.connection.Find(&tasks)
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
