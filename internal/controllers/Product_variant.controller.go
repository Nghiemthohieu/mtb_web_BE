package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductVariantsController struct {
	ProductVariantService *services.ProductVariantService
}

func NewProductVariantsController() *ProductVariantsController {
	return &ProductVariantsController{
		ProductVariantService: services.NewProductVariantService(),
	}
}

func (c *ProductVariantsController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productVariants, core, err := c.ProductVariantService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productVariants)
	}
}

func (c *ProductVariantsController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20071, err)
			return
		}
		productVariant, core, err := c.ProductVariantService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productVariant)
	}
}

func (c *ProductVariantsController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productVariant models.ProductVariant
		if err := ctx.ShouldBindJSON(&productVariant); err != nil {
			response.ErrorRespone(ctx, 400, 20071, err)
			return
		}
		core, err := c.ProductVariantService.Create(&productVariant)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productVariant)
	}
}

func (c *ProductVariantsController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productVariant models.ProductVariant
		if err := ctx.ShouldBindJSON(&productVariant); err != nil {
			response.ErrorRespone(ctx, 400, 20071, err)
			return
		}
		core, err := c.ProductVariantService.Update(&productVariant)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productVariant)
	}
}

func (c *ProductVariantsController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20071, err)
			return
		}
		core, err := c.ProductVariantService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *ProductVariantsController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20071, err)
			return
		}
		core, err := c.ProductVariantService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *ProductVariantsController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.ProductVariantService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
