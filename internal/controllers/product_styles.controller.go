package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductStylesController struct {
	ProductStylesService *services.ProductStylesService
}

func NewProductStylesController() *ProductStylesController {
	return &ProductStylesController{
		ProductStylesService: services.NewProductStylesService(),
	}
}

func (c *ProductStylesController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productStyles, core, err := c.ProductStylesService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productStyles)
	}
}
func (c *ProductStylesController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20021, err)
			return
		}
		productStyles, core, err := c.ProductStylesService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productStyles)
	}
}
func (c *ProductStylesController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productStyles models.ProductStyle
		if err := ctx.ShouldBindJSON(&productStyles); err != nil {
			response.ErrorRespone(ctx, 400, 20021, err)
			return
		}
		core, err := c.ProductStylesService.Create(&productStyles)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productStyles)
	}
}
func (c *ProductStylesController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productStyles models.ProductStyle
		if err := ctx.ShouldBindJSON(&productStyles); err != nil {
			response.ErrorRespone(ctx, 400, 20021, err)
			return
		}
		core, err := c.ProductStylesService.Update(&productStyles)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productStyles)
	}
}
func (c *ProductStylesController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20021, err)
			return
		}
		core, err := c.ProductStylesService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
func (c *ProductStylesController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20021, err)
			return
		}
		core, err := c.ProductStylesService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
func (c *ProductStylesController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.ProductStylesService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
