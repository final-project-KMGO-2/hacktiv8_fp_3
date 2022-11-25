package controller

import (
	"hacktiv8_fp_2/common"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	CreateCategory(ctx *gin.Context)
	GetCategory(ctx *gin.Context)
	PatchCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

type categoryController struct {
	categoryService service.CategoryService
	jwtService      service.JWTService
}

func NewCategoryController(cas service.CategoryService, js service.JWTService) CategoryController {
	return &categoryController{
		categoryService: cas,
		jwtService:      js,
	}
}

// CreateCategory implements CategoryController
func (ca *categoryController) CreateCategory(ctx *gin.Context) {
	var categoryCreate entity.CategoryCreate
	err := ctx.ShouldBind(&categoryCreate)
	if err != nil {
		response := common.BuildErrorResponse("Failed to bind category request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := ca.categoryService.CreateCategory(ctx.Request.Context(), categoryCreate)
	if err != nil {
		res := common.BuildErrorResponse("Failed to add category", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Accepted", result)
	ctx.JSON(http.StatusAccepted, res)
}

// GetCategory implements CategoryController
func (ca *categoryController) GetCategory(ctx *gin.Context) {

	result, err := ca.categoryService.GetCategory(ctx.Request.Context())
	if err != nil {
		response := common.BuildErrorResponse("Failed to get category", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := common.BuildResponse(true, "Accepted", result)
	ctx.JSON(http.StatusAccepted, res)
}

// PatchCategory implements CategoryController
func (ca *categoryController) PatchCategory(ctx *gin.Context) {
	var categoryPatch entity.CategoryPatch
	id := ctx.Param("id")
	categoryPatch.ID, _ = strconv.ParseUint(id, 10, 64)
	if err := ctx.ShouldBind(&categoryPatch); err != nil {
		response := common.BuildErrorResponse("Failed to bind photo request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := ca.categoryService.PatchCategory(ctx.Request.Context(), categoryPatch.ID, categoryPatch)
	if err != nil {
		response := common.BuildErrorResponse("Failed to patch category", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := common.BuildResponse(true, "Accepted", result)
	ctx.JSON(http.StatusAccepted, res)
}

// DeleteCategory implements CategoryController
func (ca *categoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryID, _ := strconv.ParseUint(id, 10, 64)
	err := ca.categoryService.DeleteCategory(ctx.Request.Context(), categoryID)
	if err != nil {
		response := common.BuildErrorResponse("Failed to delete comment", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := common.BuildResponse(true, "Category has been successfully deleted", nil)
	ctx.JSON(http.StatusAccepted, res)
}
