package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentMethodController struct {
	PaymentMethodService *services.PaymentMethodService
}

func NewPaymentMethodController() *PaymentMethodController {
	return &PaymentMethodController{
		PaymentMethodService: services.NewPaymentMethodService(),
	}
}

func (c *PaymentMethodController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paymentMethods, core, err := c.PaymentMethodService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, paymentMethods)
	}
}

func (c *PaymentMethodController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20131, err)
			return
		}
		paymentMethod, core, err := c.PaymentMethodService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, paymentMethod)
	}
}

func (c *PaymentMethodController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paymentMethod models.PaymentMethod
		if err := ctx.ShouldBindJSON(&paymentMethod); err != nil {
			response.ErrorRespone(ctx, 400, 20131, err)
			return
		}
		core, err := c.PaymentMethodService.Create(&paymentMethod)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, paymentMethod)
	}
}

func (c *PaymentMethodController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paymentMethod models.PaymentMethod
		if err := ctx.ShouldBindJSON(&paymentMethod); err != nil {
			response.ErrorRespone(ctx, 400, 20131, err)
			return
		}
		core, err := c.PaymentMethodService.Update(&paymentMethod)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, paymentMethod)
	}
}

func (c *PaymentMethodController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20131, err)
			return
		}
		core, err := c.PaymentMethodService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *PaymentMethodController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20131, err)
			return
		}
		core, err := c.PaymentMethodService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *PaymentMethodController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.PaymentMethodService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
