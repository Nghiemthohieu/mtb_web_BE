package controllers

import (
	"mtb_web/internal/dto"
	"mtb_web/internal/services"
	"mtb_web/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	CartService services.CartService
}

func NewCartController() *CartController {
	return &CartController{
		CartService: *services.NewCartService(),
	}
}

func (c *CartController) CreateRedis() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Query("session_id")
		var cart dto.RedisCartItem
		if err := ctx.ShouldBindJSON(&cart); err != nil {
			response.ErrorRespone(ctx, 400, 20161, err)
			return
		}
		err := c.CartService.CreateRedis(sessionId, &cart)
		if err != nil {
			response.ErrorRespone(ctx, 500, 20161, err)
			return
		}
		response.SuccessResponse(ctx, 200, nil)
	}
}

func (c *CartController) GetRedis() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Query("session_id")
		cart, err := c.CartService.GetRedis(sessionId)
		if err != nil {
			response.ErrorRespone(ctx, 500, 20161, err)
			return
		}
		response.SuccessResponse(ctx, 200, cart)
	}
}

func (c *CartController) SyncCartFromRedisToDb() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Query("session_id")
		userId, err := strconv.Atoi(ctx.Param("user_id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20161, err)
			return
		}
		err = c.CartService.SyncCartFromRedisToDb(sessionId, uint(userId))
		if err != nil {
			response.ErrorRespone(ctx, 500, 20161, err)
			return
		}
		response.SuccessResponse(ctx, 200, nil)
	}
}

func (c *CartController) SyncCartFromDbToRedis() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Query("session_id")
		userId, err := strconv.Atoi(ctx.Param("user_id"))
		if err != nil {
			response.ErrorRespone(ctx, 400, 20161, err)
			return
		}
		err = c.CartService.SyncCartFromDbToRedis(sessionId, uint(userId))
		if err != nil {
			response.ErrorRespone(ctx, 500, 20161, err)
			return
		}
		response.SuccessResponse(ctx, 200, nil)
	}
}