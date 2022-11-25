package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(route *gin.Engine, taskController controller.TaskController, jwtSvc service.JWTService) {
	taskRoute := route.Group("/task")
	{
		taskRoute.GET("", middleware.Authenticate(jwtSvc, "member"), taskController.GetAllTask)
		taskRoute.POST("", middleware.Authenticate(jwtSvc, "member"), taskController.AddNewTask)
		taskRoute.PATCH("/update-status/:taskId", middleware.Authenticate(jwtSvc, "member"), taskController.UpdateTaskStatus)
		taskRoute.PATCH("/update-category/:taskId", middleware.Authenticate(jwtSvc, "member"), taskController.UpdateTaskCategory)
		taskRoute.DELETE("/:taskId", middleware.Authenticate(jwtSvc, "member"), taskController.DeleteTaskById)
	}
}
