package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductMaterialsController struct {
	ProductMaterialsService *services.ProductMaterialsService
}

func NewProductMaterialsController() *ProductMaterialsController {
	return &ProductMaterialsController{
		ProductMaterialsService: services.NewProductMaterialsService(),
	}
}

func (c *ProductMaterialsController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productMaterials, core, err := c.ProductMaterialsService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productMaterials)
	}
}

func (c *ProductMaterialsController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20041, err)
			return
		}
		productMaterial, core, err := c.ProductMaterialsService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productMaterial)
	}
}

func (c *ProductMaterialsController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productMaterial models.ProductMaterial
		if err := ctx.ShouldBindJSON(&productMaterial); err != nil {
			response.ErrorRespone(ctx, 400, 20041, err)
			return
		}
		core, err := c.ProductMaterialsService.Create(&productMaterial)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productMaterial)
	}
}

func (c *ProductMaterialsController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productMaterial models.ProductMaterial
		if err := ctx.ShouldBindJSON(&productMaterial); err != nil {
			response.ErrorRespone(ctx, 400, 20041, err)
			return
		}
		core, err := c.ProductMaterialsService.Update(&productMaterial)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productMaterial)
	}
}

func (c *ProductMaterialsController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20041, err)
			return
		}
		core, err := c.ProductMaterialsService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *ProductMaterialsController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20041, err)
			return
		}
		core, err := c.ProductMaterialsService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *ProductMaterialsController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.ProductMaterialsService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
