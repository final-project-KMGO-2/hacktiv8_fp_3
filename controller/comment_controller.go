package controller

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetComment(ctx *gin.Context)
	GetCommentByID(ctx *gin.Context)
	UpdateCommentByID(ctx *gin.Context)
	DeleteCommentByID(ctx *gin.Context)
}

type commentController struct {
	commentService service.CommentService
	jwtService     service.JWTService
}

func NewCommentController(cs service.CommentService, js service.JWTService) CommentController {
	return &commentController{
		commentService: cs,
		jwtService:     js,
	}
}

// CreateComment implements CommentController
func (c *commentController) CreateComment(ctx *gin.Context) {
	var commentCreate entity.CommentCreate
	if err := ctx.ShouldBind(&commentCreate); err != nil {
		response := common.BuildErrorResponse("Failed to bind comment request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, _ := c.jwtService.GetUserIDByToken(token)

	commentCreate.UserID = uint64(userID)

	result, err := c.commentService.CreateComment(ctx.Request.Context(), commentCreate)
	if err != nil {
		response := common.BuildErrorResponse("Failed to add comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusAccepted, res)
}

// GetComment implements CommentController
func (c *commentController) GetComment(ctx *gin.Context) {

	token := ctx.MustGet("token").(string)
	userID, _ := c.jwtService.GetUserIDByToken(token)

	result, err := c.commentService.GetComment(ctx.Request.Context(), userID)
	if err != nil {
		response := common.BuildErrorResponse("Failed to get comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusAccepted, res)
}

// GetCommentByID implements CommentController
func (c *commentController) GetCommentByID(ctx *gin.Context) {
	id := ctx.Param("commentID")
	commentID, _ := strconv.ParseUint(id, 10, 64)

	result, err := c.commentService.GetCommentByID(ctx.Request.Context(), commentID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusAccepted, res)
}

// UpdateCommentByID implements CommentController
func (c *commentController) UpdateCommentByID(ctx *gin.Context) {
	var commentUpdate entity.CommentUpdate
	if err := ctx.ShouldBind(&commentUpdate); err != nil {
		response := common.BuildErrorResponse("Failed to bind photo request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, _ := c.jwtService.GetUserIDByToken(token)

	commentUpdate.UserID = uint64(userID)
	commentUpdate.ID = ctx.MustGet("commentID").(uint64)

	result, err := c.commentService.UpdateCommentByID(ctx.Request.Context(), commentUpdate.ID, commentUpdate)
	if err != nil {
		response := common.BuildErrorResponse("Failed to update comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusAccepted, res)
}

// DeleteCommentByID implements CommentController
func (c *commentController) DeleteCommentByID(ctx *gin.Context) {
	commentID := ctx.MustGet("commentID").(uint64)
	err := c.commentService.DeleteCommentByID(ctx.Request.Context(), commentID)
	if err != nil {
		response := common.BuildErrorResponse("Failed to delete comment", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := common.BuildResponse(true, "Your comment has been successfully deleted", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
