package controller

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	authService service.AuthService
	jwtService  service.JWTService
}

func NewUserController(us service.UserService, as service.AuthService, js service.JWTService) UserController {
	return &userController{
		userService: us,
		authService: as,
		jwtService:  js,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user entity.UserRegister
	errBind := ctx.ShouldBind(&user)

	if errBind != nil {
		response := common.BuildErrorResponse("Failed to process request", errBind.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isDuplicateEmail, _ := c.authService.CheckEmailDuplicate(ctx.Request.Context(), user.Email)
	if isDuplicateEmail {
		response := common.BuildErrorResponse("Failed to process request", "Duplicate Email", common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	isDuplicateUsername, _ := c.authService.CheckUsernameDuplicate(ctx.Request.Context(), user.Username)
	if isDuplicateUsername {
		response := common.BuildErrorResponse("Failed to process request", "Duplicate Username", common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	createdUser, err := c.userService.CreateUser(ctx.Request.Context(), user)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	userId := strconv.FormatUint(uint64(createdUser.ID), 10)
	token := c.jwtService.GenerateToken(userId)
	response := common.BuildResponse(true, "OK", token)
	ctx.JSON(http.StatusCreated, response)
}

func (c *userController) Login(ctx *gin.Context) {
	var userLogin entity.UserLogin
	if errBind := ctx.ShouldBind(&userLogin); errBind != nil {
		response := common.BuildErrorResponse("Failed to process request", errBind.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult, _ := c.authService.VerifyCredential(ctx.Request.Context(), userLogin.Email, userLogin.Password)
	if !authResult {
		response := common.BuildErrorResponse("Error Logging in", "Invalid Credentials", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, err := c.userService.GetUserByEmail(ctx.Request.Context(), userLogin.Email)
	if err != nil {
		response := common.BuildErrorResponse("Failed to process request", err.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := strconv.FormatUint(uint64(user.ID), 10)
	generatedToken := c.jwtService.GenerateToken(userId)
	response := common.BuildResponse(true, "OK", generatedToken)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var userUpdate entity.UserUpdate
	errBind := ctx.ShouldBind(&userUpdate)

	if errBind != nil {
		response := common.BuildErrorResponse("Failed to process request", errBind.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, _ := c.jwtService.GetUserIDByToken(token)
	userUpdate.ID = uint64(userID)
	result, err := c.userService.UpdateUser(ctx.Request.Context(), userUpdate)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update user", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, _ := c.jwtService.GetUserIDByToken(token)
	err := c.userService.DeleteUser(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to delete user", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Your account has been successfully deleted", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
