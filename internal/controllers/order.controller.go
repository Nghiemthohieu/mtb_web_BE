package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService *services.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		OrderService: services.NewOrderService(),
	}
}

func (c *OrderController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orders, core, err := c.OrderService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, orders)
	}
}

func (c *OrderController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		order, core, err := c.OrderService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, order)
	}
}

func (c *OrderController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var order models.Order
		if err := ctx.ShouldBindJSON(&order); err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		core, err := c.OrderService.Create(&order)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, order)
	}
}

func (c *OrderController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var order models.Order
		if err := ctx.ShouldBindJSON(&order); err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		core, err := c.OrderService.Update(&order)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, order)
	}
}

func (c *OrderController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		core, err := c.OrderService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *OrderController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		core, err := c.OrderService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *OrderController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.OrderService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *OrderController) GetByCustomerID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		orders, core, err := c.OrderService.GetByCustomerID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, orders)
	}
}
