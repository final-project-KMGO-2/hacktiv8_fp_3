package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine, commentController controller.CommentController, commentService service.CommentService, jwtService service.JWTService) {
	commentRoutes := router.Group("/comments", middleware.Authenticate(jwtService))
	{
		commentRoutes.POST("", commentController.CreateComment)
		commentRoutes.GET("", commentController.GetComment)
		commentRoutes.GET("/:commentID", commentController.GetCommentByID)
		commentRoutes.PUT("/:commentID", middleware.CommentAuthorization(jwtService, commentService), commentController.UpdateCommentByID)
		commentRoutes.DELETE("/:commentID", middleware.CommentAuthorization(jwtService, commentService), commentController.DeleteCommentByID)
	}
}
