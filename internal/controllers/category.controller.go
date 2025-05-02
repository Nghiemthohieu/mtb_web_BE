package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	CategoryService *services.CategoryService
}

func NewCategoriesController() *CategoriesController {
	return &CategoriesController{
		CategoryService: services.NewCategoryService(),
	}
}

func (c *CategoriesController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categories, core, err := c.CategoryService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, categories)
	}
}

func (c *CategoriesController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20061, err)
			return
		}
		category, core, err := c.CategoryService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, category)
	}
}

func (c *CategoriesController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var category models.Category
		if err := ctx.ShouldBindJSON(&category); err != nil {
			response.ErrorRespone(ctx, 400, 20061, err)
			return
		}
		core, err := c.CategoryService.Create(&category)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, category)
	}
}

func (c *CategoriesController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var category models.Category
		if err := ctx.ShouldBindJSON(&category); err != nil {
			response.ErrorRespone(ctx, 400, 20061, err)
			return
		}
		core, err := c.CategoryService.Update(&category)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, category)
	}
}

func (c *CategoriesController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20061, err)
			return
		}
		core, err := c.CategoryService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *CategoriesController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20061, err)
			return
		}
		core, err := c.CategoryService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *CategoriesController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.CategoryService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
