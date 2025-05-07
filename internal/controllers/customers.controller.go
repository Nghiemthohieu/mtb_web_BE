package controllers

import (
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	CustomerService *services.CustomerService
}

func NewCustomerController() *CustomerController {
	return &CustomerController{
		CustomerService: services.NewCustomerService(),
	}
}

func (c *CustomerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customers, core, err := c.CustomerService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, customers)
	}
}

func (c *CustomerController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20111, err)
			return
		}
		customer, core, err := c.CustomerService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, customer)
	}
}

func (c *CustomerController) GetByEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.Param("email")
		customer, core, err := c.CustomerService.GetByEmail(email)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, customer)
	}
}

func (c *CustomerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20111, err)
			return
		}
		core, err := c.CustomerService.Delete(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
