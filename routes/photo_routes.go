package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func PhotoRoutes(router *gin.Engine, photoController controller.PhotoController, photoService service.PhotoService, jwtService service.JWTService) {
	photoRoutes := router.Group("/photos", middleware.Authenticate(jwtService))
	{
		photoRoutes.GET("/", photoController.GetPhotos)
		photoRoutes.POST("/", photoController.CreatePhoto)
		photoRoutes.GET("/:photoID", photoController.GetPhotoByID)
		photoRoutes.PUT("/:photoID", middleware.PhotoAuthorization(jwtService, photoService), photoController.UpdatePhoto)
		photoRoutes.DELETE("/:photoID", middleware.PhotoAuthorization(jwtService, photoService), photoController.DeletePhoto)
	}
}
