package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductSizeController struct {
	ProductSizeService *services.ProductSizeService
}

func NewProductSizeController() *ProductSizeController {
	return &ProductSizeController{
		ProductSizeService: services.NewProductSizeService(),
	}
}

func (c *ProductSizeController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productSizes, core, err := c.ProductSizeService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productSizes)
	}
}
func (c *ProductSizeController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20011, err)
			return
		}
		productSize, core, err := c.ProductSizeService.GetByID(uint(id))
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productSize)
	}
}
func (c *ProductSizeController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productSize models.ProductSize
		if err := ctx.ShouldBindJSON(&productSize); err != nil {
			response.ErrorRespone(ctx, 400, 20011, err)
			return
		}
		core, err := c.ProductSizeService.Create(&productSize)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, "Tạo thành công")
	}
}
func (c *ProductSizeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productSize models.ProductSize
		if err := ctx.ShouldBindJSON(&productSize); err != nil {
			response.ErrorRespone(ctx, 400, 20011, err)
			return
		}
		core, err := c.ProductSizeService.Update(&productSize)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, productSize)
	}
}
func (c *ProductSizeController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20011, err)
			return
		}
		core, err := c.ProductSizeService.Delete(uint(id))
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
func (c *ProductSizeController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []uint
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20011, err)
			return
		}
		core, err := c.ProductSizeService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
func (c *ProductSizeController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.ProductSizeService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
