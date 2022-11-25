package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(route *gin.Engine, taskController controller.TaskController, jwtSvc service.JWTService){
	taskRoute := route.Group("/task")
	{
		taskRoute.GET("", taskController.GetAllTask)
		taskRoute.POST("", taskController.AddNewTask)
		taskRoute.PATCH("/update-status/:taskId", taskController.UpdateTaskStatus)
		taskRoute.PATCH("/update-category/:taskId", taskController.UpdateTaskCategory)
		taskRoute.DELETE("/:taskId", taskController.DeleteTaskById)
	}
}