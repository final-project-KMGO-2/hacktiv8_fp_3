package routes

import "github.com/gin-gonic/gin"

func TaskRoutes(route *gin.Engine){
	taskRoute := route.Group("/task")
	{
		taskRoute.GET("")
		taskRoute.POST("")
		taskRoute.PATCH("/update-status/:id")
		taskRoute.PATCH("/update-category/:id")
		taskRoute.DELETE("/:id")
	}
}