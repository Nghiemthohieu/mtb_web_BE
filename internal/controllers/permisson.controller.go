package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	PermissionService *services.PermissionService
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		PermissionService: services.NewPermissionService(),
	}
}

func (c *PermissionController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		permissions, core, err := c.PermissionService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, permissions)
	}
}

func (c *PermissionController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20121, err)
			return
		}
		permission, core, err := c.PermissionService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, permission)
	}
}

func (c *PermissionController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var permission models.Permission
		if err := ctx.ShouldBindJSON(&permission); err != nil {
			response.ErrorRespone(ctx, 400, 20121, err)
			return
		}
		core, err := c.PermissionService.Create(&permission)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, permission)
	}
}

func (c *PermissionController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var permission models.Permission
		if err := ctx.ShouldBindJSON(&permission); err != nil {
			response.ErrorRespone(ctx, 400, 20121, err)
			return
		}
		core, err := c.PermissionService.Update(&permission)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, permission)
	}
}

func (c *PermissionController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20121, err)
			return
		}
		core, err := c.PermissionService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *PermissionController) DeleteByIDs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ids []int
		if err := ctx.ShouldBindJSON(&ids); err != nil {
			response.ErrorRespone(ctx, 400, 20121, err)
			return
		}
		core, err := c.PermissionService.DeleteByIDs(ids)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}

func (c *PermissionController) DeleteAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core, err := c.PermissionService.DeleteAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
