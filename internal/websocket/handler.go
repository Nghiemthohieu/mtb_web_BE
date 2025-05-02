package websocket

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.Query("user_id")
		userID, _ := strconv.Atoi(userIDStr)

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		client := &Client{
			ID:   uint(userID),
			Conn: conn,
			Send: make(chan []byte),
		}

		hub.Register <- client
		go client.WritePump()
	}
}
