package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(route *gin.Engine, taskController controller.TaskController, jwtSvc service.JWTService, taskService service.TaskService) {
	taskRoute := route.Group("/tasks", middleware.Authenticate(jwtSvc, "member"))
	{
		taskRoute.GET("", taskController.GetAllTask)
		taskRoute.POST("", taskController.AddNewTask)
		taskRoute.PUT("/:taskId", middleware.TaskAuthorization(jwtSvc, taskService), taskController.UpdateTask)
		taskRoute.PATCH("/update-status/:taskId", middleware.TaskAuthorization(jwtSvc, taskService), taskController.UpdateTaskStatus)
		taskRoute.PATCH("/update-category/:taskId", middleware.TaskAuthorization(jwtSvc, taskService), taskController.UpdateTaskCategory)
		taskRoute.DELETE("/:taskId", middleware.TaskAuthorization(jwtSvc, taskService), taskController.DeleteTaskById)
	}
}
