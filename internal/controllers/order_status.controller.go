package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderStatusController struct {
	OrderStatusService *services.OrderStatusService
}

func NewOrderStatusController() *OrderStatusController {
	return &OrderStatusController{
		OrderStatusService: services.NewOrderStatusService(),
	}
}

func (c *OrderStatusController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderStatuses, core, err := c.OrderStatusService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, orderStatuses)
	}
}

func (c *OrderStatusController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20081, err)
			return
		}
		orderStatus, core, err := c.OrderStatusService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, orderStatus)
	}
}

func (c *OrderStatusController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orderStatus models.OrderStatus
		if err := ctx.ShouldBindJSON(&orderStatus); err != nil {
			response.ErrorRespone(ctx, 400, 20081, err)
			return
		}
		core, err := c.OrderStatusService.Create(&orderStatus)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, orderStatus)
	}
}

func (c *OrderStatusController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orderStatus models.OrderStatus
		if err := ctx.ShouldBindJSON(&orderStatus); err != nil {
			response.ErrorRespone(ctx, 400, 20081, err)
			return
		}
		core, err := c.OrderStatusService.Update(&orderStatus)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, orderStatus)
	}
}

func (c *OrderStatusController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20081, err)
			return
		}
		core, err := c.OrderStatusService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *OrderStatusController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20081, err)
			return
		}
		core, err := c.OrderStatusService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *OrderStatusController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.OrderStatusService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
