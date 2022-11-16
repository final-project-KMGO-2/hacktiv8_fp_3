package middleware

import (
	"fmt"
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := common.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := common.BuildErrorResponse("Invalid Token", "Token is invalid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if !token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			log.Println(err)
			response := common.BuildErrorResponse("Access denied", "You dont have access", nil)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		c.Set("token", authHeader)
		c.Next()
	}
}
