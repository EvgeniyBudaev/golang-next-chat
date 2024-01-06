package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
	Conn     *websocket.Conn
	Message  chan *Message
}

type ClientResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}
		h.Broadcast <- msg
	}
}
