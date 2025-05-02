package controllers

import (
	"fmt"
	"mtb_web/internal/dto"
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	ProductService *services.ProductService
}

func NewProductsController() *ProductsController {
	return &ProductsController{
		ProductService: services.NewProductService(),
	}
}

func (c *ProductsController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, core, err := c.ProductService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, 20001, products)
	}
}

func (c *ProductsController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20051, err)
			return
		}
		product, core, err := c.ProductService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, 20001, product)
	}
}

func (c *ProductsController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product models.Product
		if err := ctx.ShouldBindJSON(&product); err != nil {
			response.ErrorRespone(ctx, 400, 20051, err)
			return
		}
		core, err := c.ProductService.Create(&product)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, product)
	}
}

func (c *ProductsController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product models.Product
		if err := ctx.ShouldBindJSON(&product); err != nil {
			response.ErrorRespone(ctx, 400, 20051, err)
			return
		}
		core, err := c.ProductService.Update(&product)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, product)
	}
}

func (c *ProductsController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20051, err)
			return
		}
		core, err := c.ProductService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *ProductsController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20051, err)
			return
		}
		core, err := c.ProductService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *ProductsController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.ProductService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *ProductsController) GetByCategoryID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20051, err)
			return
		}
		products, core, err := c.ProductService.GetByCategoryID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, 20001, products)
	}
}

func (c *ProductsController) GetByCaregoryIDAndColorID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryID, err := strconv.Atoi(ctx.Param("category_id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20051, err)
			return
		}
		colorID, err := strconv.Atoi(ctx.Param("color_id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20052, err)
			return
		}
		products, core, err := c.ProductService.GetByCaregoryIDAndColorID(categoryID, colorID)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, 20001, products)
	}
}

func (c *ProductsController) GetByCaregoryIDAndMaterialID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryID, err := strconv.Atoi(ctx.Param("category_id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20051, err)
			return
		}
		materialID, err := strconv.Atoi(ctx.Param("material_id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20052, err)
			return
		}
		products, core, err := c.ProductService.GetByCaregoryIDAndMaterialID(categoryID, materialID)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, 20001, products)
	}
}

func (c *ProductsController) GetByCategoryIDAndSizeIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.CategorySizeRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.ErrorRespone(ctx, 400, 20051, fmt.Errorf("invalid JSON body: %v", err))
			return
		}

		if req.CategoryId <= 0 {
			response.ErrorRespone(ctx, 400, 20052, fmt.Errorf("category_id must be a positive integer"))
			return
		}

		if len(req.SizeIds) == 0 {
			response.ErrorRespone(ctx, 400, 20053, fmt.Errorf("size_ids cannot be empty"))
			return
		}

		products, core, err := c.ProductService.GetByCaregoryIDAndSizeIDs(req.CategoryId, req.SizeIds)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}

		response.SuccessResponse(ctx, 20001, products)
	}
}

func (c *ProductsController) GetByCaregoryIDAndStyleID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryID, err := strconv.Atoi(ctx.Param("category_id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20051, err)
			return
		}
		styleID, err := strconv.Atoi(ctx.Param("style_id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20052, err)
			return
		}
		products, core, err := c.ProductService.GetByCaregoryIDAndStyleID(categoryID, styleID)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, 20001, products)
	}
}
