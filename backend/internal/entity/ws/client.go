package ws

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	ID       int64  `json:"id"`
	RoomID   int64  `json:"roomId"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Conn     *websocket.Conn
	Message  chan *Message
}

type ClientResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type Message struct {
	ID       int64  `json:"id"`
	RoomID   int64  `json:"roomId"`
	ClientID int64  `json:"clientId"`
	Content  string `json:"content"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		fmt.Println("[WriteMessage c] ", c)
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
		fmt.Println("[ReadMessage m] ", string(m))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg := &Message{
			ID:       c.ID,
			RoomID:   c.RoomID,
			ClientID: c.ID,
			Content:  string(m),
		}
		h.Broadcast <- msg
	}
}
