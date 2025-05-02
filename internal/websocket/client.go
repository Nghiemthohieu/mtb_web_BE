package websocket

import (
    "github.com/gorilla/websocket"
)

type Client struct {
    ID   uint
    Conn *websocket.Conn
    Send chan []byte
}

func (c *Client) WritePump() {
    for {
        msg, ok := <-c.Send
        if !ok {
            return
        }
        c.Conn.WriteMessage(websocket.TextMessage, msg)
    }
}
