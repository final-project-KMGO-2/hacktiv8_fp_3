package controller

import "github.com/gin-gonic/gin"

type TaskController interface {
	AddNewTask(ctx *gin.Context)
	GetAllTask(ctx *gin.Context)
	UpdateTaskStatus(ctx *gin.Context)
	UpdateTaskCategory(ctx *gin.Context)
	DeleteTaskById(ctx *gin.Context)
}