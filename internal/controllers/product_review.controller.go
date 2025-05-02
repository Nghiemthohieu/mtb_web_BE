package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductReviewsController struct {
	ProductReviewsService *services.ProductReviewService
}

func NewProductReviewsController() *ProductReviewsController {
	return &ProductReviewsController{
		ProductReviewsService: services.NewProductReviewService(),
	}
}

func (c *ProductReviewsController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reviews, core, err := c.ProductReviewsService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, reviews)
	}
}

func (c *ProductReviewsController) GetByProductID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productID, err := strconv.Atoi(ctx.Param("product_id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20021, err)
			return
		}
		reviews, core, err := c.ProductReviewsService.GetByProductID(uint(productID))
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, reviews)
	}
}

func (c *ProductReviewsController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var review models.ProductReview
		if err := ctx.ShouldBindJSON(&review); err != nil {
			response.ErrorRespone(ctx, 400, 20021, err)
			return
		}
		core, err := c.ProductReviewsService.Create(&review)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, review)
	}
}

func (c *ProductReviewsController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var review models.ProductReview
		if err := ctx.ShouldBindJSON(&review); err != nil {
			response.ErrorRespone(ctx, 400, 20021, err)
			return
		}
		core, err := c.ProductReviewsService.Update(&review)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, review)
	}
}

func (c *ProductReviewsController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20021, err)
			return
		}
		core, err := c.ProductReviewsService.DeleteById(uint(id))
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
