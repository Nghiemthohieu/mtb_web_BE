package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductColorController struct {
	ProductColorService *services.ProductColorService
}

func NewProductColorController() *ProductColorController {
	return &ProductColorController{
		ProductColorService: services.NewProductColorService(),
	}
}

func (c *ProductColorController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productColor models.ProductColor
		if err := ctx.ShouldBindJSON(&productColor); err != nil {
			response.ErrorRespone(ctx, 400, 20031, err)
			return
		}
		code, err := c.ProductColorService.Create(&productColor)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product color created successfully")
	}
}

func (c *ProductColorController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productColor models.ProductColor
		if err := ctx.ShouldBindJSON(&productColor); err != nil {
			response.ErrorRespone(ctx, 400, 20031, err)
			return
		}
		code, err := c.ProductColorService.Update(&productColor)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product color updated successfully")
	}
}
func (c *ProductColorController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20031, err)
			return
		}
		productColor, code, err := c.ProductColorService.GetById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, productColor)
	}
}

func (c *ProductColorController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productColors, code, err := c.ProductColorService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, productColors)
	}
}
func (c *ProductColorController) DeleteById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20031, err)
			return
		}
		code, err := c.ProductColorService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product color deleted successfully")
	}
}
func (c *ProductColorController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20031, err)
			return
		}
		code, err := c.ProductColorService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "Product colors deleted successfully")
	}
}
func (c *ProductColorController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code, err := c.ProductColorService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, code, err)
			return
		}
		response.SuccessResponse(ctx, code, "All product colors deleted successfully")
	}
}
