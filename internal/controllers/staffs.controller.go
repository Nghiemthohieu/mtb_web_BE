package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StaffsController struct {
	StaffService services.StaffService
}

func NewStaffsController() *StaffsController {
	return &StaffsController{
		StaffService: *services.NewStaffService(),
	}
}

func (c *StaffsController) GetAllStaffs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		staffs, core, err := c.StaffService.GetAllStaff()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, staffs)
	}
}

func (c *StaffsController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		staff, core, err := c.StaffService.GetStaffById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, staff)
	}
}

func (c *StaffsController) GetByEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gmail := ctx.Param("email")
		staff, core, err := c.StaffService.GetStaffByEmail(gmail)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, staff)
	}
}

func (c *StaffsController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var staff models.User
		if err := ctx.ShouldBindJSON(&staff); err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		core, err := c.StaffService.Create(&staff)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, staff)
	}
}

func (c *StaffsController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var staff models.User
		if err := ctx.ShouldBindJSON(&staff); err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		core, err := c.StaffService.UpdateFullStaff(&staff)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, staff)
	}
}

func (c *StaffsController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20091, err)
			return
		}
		core, err := c.StaffService.Delete(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
