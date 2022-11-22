package controller

import (
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	AddNewTask(ctx *gin.Context)
	GetAllTask(ctx *gin.Context)
	UpdateTaskStatus(ctx *gin.Context)
	UpdateTaskCategory(ctx *gin.Context)
	DeleteTaskById(ctx *gin.Context)
}

type taskController struct {
	taskService service.TaskService
}

func NewTaskService(ts service.TaskService) TaskController {
	return taskController{taskService: ts}
}

func (tc taskController) AddNewTask(ctx *gin.Context)
func (tc taskController) GetAllTask(ctx *gin.Context)
func (tc taskController) UpdateTaskStatus(ctx *gin.Context)
func (tc taskController) UpdateTaskCategory(ctx *gin.Context)
func (tc taskController) DeleteTaskById(ctx *gin.Context)