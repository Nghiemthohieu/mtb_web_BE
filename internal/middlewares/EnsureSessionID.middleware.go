package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func EnsureSessionID() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("session_id")
		if err != nil || sessionId == "" {
			// Nếu không có session_id thì tạo mới
			sessionId = uuid.New().String()
			c.SetCookie("session_id", sessionId, 3600*24*30, "/", "localhost", false, true)
		}

		// Đặt session_id vào context để dùng trong handler/service
		c.Set("session_id", sessionId)
		c.Next()
	}
}
