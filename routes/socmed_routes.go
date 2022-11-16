package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func SocMedRoutes(router *gin.Engine, socmedController controller.SocmedController, socmedService service.SocmedService, jwtService service.JWTService) {
	smRoutes := router.Group("/socialmedias", middleware.Authenticate(jwtService))
	{
		smRoutes.POST("", socmedController.PostSocmed)
		smRoutes.GET("", socmedController.GetSocmed)
		smRoutes.PUT("/:socialMediaId", middleware.SocmedAuthorization(jwtService, socmedService), socmedController.UpdateSocmedById)
		smRoutes.DELETE("/:socialMediaId", middleware.SocmedAuthorization(jwtService, socmedService), socmedController.DeleteSocmedById)
	}
}
