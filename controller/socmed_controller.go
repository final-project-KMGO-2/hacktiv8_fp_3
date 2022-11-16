package controller

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocmedController interface {
	PostSocmed(ctx *gin.Context)
	GetSocmed(ctx *gin.Context)
	UpdateSocmedById(ctx *gin.Context)
	DeleteSocmedById(ctx *gin.Context)
}

type socmedController struct {
	userService   service.UserService
	jwtService    service.JWTService
	socmedService service.SocmedService
}

func NewSocmedController(us service.UserService, ss service.SocmedService, js service.JWTService) SocmedController {
	return &socmedController{
		userService:   us,
		socmedService: ss,
		jwtService:    js,
	}
}

func (sc *socmedController) PostSocmed(ctx *gin.Context) {
	var socmedCreate entity.SocialMediaCreate
	err := ctx.ShouldBind(&socmedCreate)
	if err != nil {
		response := common.BuildErrorResponse("invalid input", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
	userId, err := sc.jwtService.GetUserIDByToken(token)
	socmedCreate.UserID = uint64(userId)

	if err != nil {
		response := common.BuildErrorResponse("invalid token", err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	result, err := sc.socmedService.AddNewSocmed(ctx.Request.Context(), socmedCreate) // implement func ...

	if err != nil {
		response := common.BuildErrorResponse("invalid input", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "Created", result)
	ctx.JSON(http.StatusCreated, response)
}

func (sc *socmedController) GetSocmed(ctx *gin.Context) {
	data, err := sc.socmedService.GetSocmedInfo(ctx)

	if err != nil {
		response := common.BuildErrorResponse("Something went wrong, failed to get social medias", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := common.BuildResponse(true, "OK", data)
	ctx.JSON(http.StatusOK, response)
}
func (sc *socmedController) UpdateSocmedById(ctx *gin.Context) {

	var socmedUpdate entity.SocialMediaUpdate
	err := ctx.ShouldBind(&socmedUpdate)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong, failed to update the social media", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	token := ctx.MustGet("token").(string)
	userId, err := sc.jwtService.GetUserIDByToken(token)

	if err != nil {
		response := common.BuildErrorResponse("Something went wrong, failed to update the social media", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	socmedUpdate.UserID = uint64(userId)
	socmedUpdate.ID = ctx.MustGet("socmedID").(uint64)

	data, err := sc.socmedService.UpdateSocmed(ctx, socmedUpdate)

	if err != nil {
		response := common.BuildErrorResponse("Something went wrong, failed to update the social media", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := common.BuildResponse(true, "OK", data)
	ctx.JSON(http.StatusOK, response)
}

func (sc *socmedController) DeleteSocmedById(ctx *gin.Context) {
	id := ctx.MustGet("socmedID").(uint64)
	err := sc.socmedService.DeleteSocmed(ctx.Request.Context(), id)
	if err != nil {
		response := common.BuildErrorResponse("Something went wrong, failed to delete the social media", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := common.BuildResponse(true, "Your social media has been successfully deleted", common.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
