package controller

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	AddNewTask(ctx *gin.Context)
	GetAllTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	UpdateTaskStatus(ctx *gin.Context)
	UpdateTaskCategory(ctx *gin.Context)
	DeleteTaskById(ctx *gin.Context)
}

type taskController struct {
	taskService service.TaskService
	jwtService  service.JWTService
}

func NewTaskService(ts service.TaskService, jt service.JWTService) TaskController {
	return taskController{taskService: ts, jwtService: jt}
}

func (tc taskController) AddNewTask(ctx *gin.Context) {
	var taskCreate entity.TaskCreate
	err := ctx.ShouldBind(&taskCreate)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	token := ctx.MustGet("token").(string)
	userId, err := tc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
	taskCreate.UserID = int(userId)
	data, err := tc.taskService.CreateNewTask(ctx.Request.Context(), taskCreate)

	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "Task Created", data)
	ctx.JSON(http.StatusCreated, response)
}

func (tc taskController) GetAllTask(ctx *gin.Context) {
	data, err := tc.taskService.GetTasks(ctx.Request.Context())
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := common.BuildResponse(true, "OK", data)
	ctx.JSON(http.StatusOK, response)
}

func (tc taskController) UpdateTask(ctx *gin.Context) {
	var taskUpdate entity.TaskUpdate
	err := ctx.ShouldBind(&taskUpdate)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
	userId, err := tc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	data, err := tc.taskService.UpdateTask(ctx, taskUpdate, int(userId))
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "OK", data)
	ctx.JSON(http.StatusOK, response)
}

func (tc taskController) UpdateTaskStatus(ctx *gin.Context) {
	var taskStatusReq entity.TaskStatusModifier
	err := ctx.ShouldBind(&taskStatusReq)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	token := ctx.MustGet("token").(string)

	userId, err := tc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	data, err := tc.taskService.ChangeTaskStatus(ctx, taskStatusReq, int(userId))
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "OK", data)
	ctx.JSON(http.StatusOK, response)
}

func (tc taskController) UpdateTaskCategory(ctx *gin.Context) {
	var taskCategoryReq entity.TaskCategoryModifier
	err := ctx.ShouldBind(&taskCategoryReq)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	token := ctx.MustGet("token").(string)

	userId, err := tc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	data, err := tc.taskService.ChangeTaskCategory(ctx, taskCategoryReq, int(userId))
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "OK", data)
	ctx.JSON(http.StatusOK, response)
}

func (tc taskController) DeleteTaskById(ctx *gin.Context) {
	id := ctx.MustGet("taskId").(int)
	err := tc.taskService.RemoveTask(ctx.Request.Context(), id)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "Task Deleted", common.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
