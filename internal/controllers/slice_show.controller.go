package controllers

import (
	"mtb_web/internal/models"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SlideShowController struct {
	SlideShowService *services.SlideShowService
}

func NewSlideShowController() *SlideShowController {
	return &SlideShowController{
		SlideShowService: services.NewSlideShowService(),
	}
}

func (c *SlideShowController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		slides, core, err := c.SlideShowService.GetAll()
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, slides)
	}
}

func (c *SlideShowController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20171, err)
			return
		}
		slide, core, err := c.SlideShowService.GetByID(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, slide)
	}
}

func (c *SlideShowController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var slide models.SlideShow
		if err := ctx.ShouldBindJSON(&slide); err != nil {
			response.ErrorRespone(ctx, 400, 20171, err)
			return
		}
		core, err := c.SlideShowService.Create(&slide)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, slide)
	}
}

func (c *SlideShowController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var slide models.SlideShow
		if err := ctx.ShouldBindJSON(&slide); err != nil {
			response.ErrorRespone(ctx, 400, 20171, err)
			return
		}
		core, err := c.SlideShowService.Update(&slide)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, slide)
	}
}

func (c *SlideShowController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20171, err)
			return
		}
		core, err := c.SlideShowService.DeleteById(id)
		if err != nil {
			response.ErrorRespone(ctx, 500, core, err)
			return
		}
		response.SuccessResponse(ctx, core, nil)
	}
}
