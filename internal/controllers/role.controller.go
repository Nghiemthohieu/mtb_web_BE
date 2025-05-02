package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	RoleService *services.RoleService
}

func NewRoleController() *RoleController {
	return &RoleController{
		RoleService: services.NewRoleService(),
	}
}

func (c *RoleController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roles, core, err := c.RoleService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, roles)
	}
}

func (c *RoleController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20111, err)
			return
		}
		role, core, err := c.RoleService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, role)
	}
}

func (c *RoleController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var role models.Role
		if err := ctx.ShouldBindJSON(&role); err != nil {
			response.ErrorRespone(ctx, 400, 20111, err)
			return
		}
		core, err := c.RoleService.Create(&role)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, role)
	}
}

func (c *RoleController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var role models.Role
		if err := ctx.ShouldBindJSON(&role); err != nil {
			response.ErrorRespone(ctx, 400, 20111, err)
			return
		}
		core, err := c.RoleService.Update(&role)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, role)
	}
}

func (c *RoleController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20111, err)
			return
		}
		core, err := c.RoleService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *RoleController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20111, err)
			return
		}
		core, err := c.RoleService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *RoleController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.RoleService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
