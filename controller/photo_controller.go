package controller

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
	GetPhotos(ctx *gin.Context)
	GetPhotoByID(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoController struct {
	photoService service.PhotoService
	jwtService   service.JWTService
}

func NewPhotoController(ps service.PhotoService, js service.JWTService) PhotoController {
	return &photoController{
		photoService: ps,
		jwtService:   js,
	}
}

func (c *photoController) CreatePhoto(ctx *gin.Context) {
	var photo entity.PhotoCreate
	if err := ctx.ShouldBind(&photo); err != nil {
		res := common.BuildErrorResponse("Failed to bind photo request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, _ := c.jwtService.GetUserIDByToken(token)

	photo.UserID = uint64(userID)

	result, err := c.photoService.CreatePhoto(ctx.Request.Context(), photo)
	if err != nil {
		res := common.BuildErrorResponse("Failed to insert photo", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *photoController) GetPhotos(ctx *gin.Context) {
	result, err := c.photoService.GetPhotos(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get photos", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *photoController) GetPhotoByID(ctx *gin.Context) {
	id := ctx.Param("photoID")
	photoID, _ := strconv.ParseUint(id, 10, 64)

	result, err := c.photoService.GetPhotoByID(ctx.Request.Context(), photoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get photo", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *photoController) UpdatePhoto(ctx *gin.Context) {
	var photo entity.PhotoUpdate
	if err := ctx.ShouldBind(&photo); err != nil {
		res := common.BuildErrorResponse("Failed to bind photo request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, _ := c.jwtService.GetUserIDByToken(token)

	photo.UserID = uint64(userID)
	photo.ID = ctx.MustGet("photoID").(uint64)

	result, err := c.photoService.UpdatePhoto(ctx.Request.Context(), photo)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update photo", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *photoController) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.MustGet("photoID").(uint64)
	err := c.photoService.DeletePhoto(ctx.Request.Context(), photoID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to delete photo", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Your photo has been successfully deleted", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
