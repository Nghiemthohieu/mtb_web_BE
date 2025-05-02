package websocket

import "sync"

type Hub struct {
	Clients    map[uint]*Client
	Register   chan *Client
	Unregister chan *Client
	mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[uint]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client.ID] = client
			h.mu.Unlock()

		case client := <-h.Unregister:
			h.mu.Lock()
			delete(h.Clients, client.ID)
			h.mu.Unlock()
			client.Conn.Close()
		}
	}
}

func (h *Hub) SendToUser(userID uint, message []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if client, ok := h.Clients[userID]; ok {
		client.Send <- message
	}
}
