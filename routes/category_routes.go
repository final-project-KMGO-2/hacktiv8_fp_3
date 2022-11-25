package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine, cateogoryController controller.CategoryController, jwtService service.JWTService) {
	categoryRoutes := router.Group("/categories", middleware.Authenticate(jwtService, "admin"))
	{
		categoryRoutes.GET("", cateogoryController.GetCategory)
		categoryRoutes.POST("", cateogoryController.CreateCategory)
		categoryRoutes.PATCH("/:id", cateogoryController.PatchCategory)
		categoryRoutes.DELETE("/:id", cateogoryController.DeleteCategory)
	}
}
