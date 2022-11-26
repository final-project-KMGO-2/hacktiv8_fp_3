package middleware

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TaskAuthorization(jwtService service.JWTService, taskService service.TaskService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		taskID, _ := strconv.ParseInt(ctx.Param("taskId"), 10, 64)

		task, err := taskService.GetTaskByID(ctx.Request.Context(), int(taskID))
		if err != nil {
			response := common.BuildErrorResponse("Failed to process request", "task does not exist", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token := ctx.MustGet("token").(string)
		userID, _ := jwtService.GetUserIDByToken(token)

		if task.UserID != int(userID) {
			response := common.BuildErrorResponse("Failed to process request", "You are not authorized to access this data", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Set("taskID", int(taskID))
		ctx.Next()
	}
}
