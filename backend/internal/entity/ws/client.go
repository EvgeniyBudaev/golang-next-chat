package ws

import (
	"encoding/json"
	"fmt"
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"go.uber.org/zap"
	"log"
	"time"

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

type MessageType string

const (
	SystemMessage   MessageType = "sys"
	ReceivedMessage MessageType = "recv"
	SelfMessage     MessageType = "self"
	NoneMessage     MessageType = "none"
)

type Message struct {
	ID        int64       `json:"id"`
	RoomID    int64       `json:"roomId"`
	UserID    string      `json:"userId"`
	Type      MessageType `json:"type"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	IsDeleted bool        `json:"isDeleted"`
	IsEdited  bool        `json:"isEdited"`
	Content   string      `json:"content"`
}

type ResponseMessage struct {
	ID        int64                                   `json:"id"`
	RoomID    int64                                   `json:"roomId"`
	UserID    string                                  `json:"userId"`
	Type      MessageType                             `json:"type"`
	CreatedAt time.Time                               `json:"createdAt"`
	UpdatedAt time.Time                               `json:"updatedAt"`
	IsDeleted bool                                    `json:"isDeleted"`
	IsEdited  bool                                    `json:"isEdited"`
	Profile   *profileEntity.ResponseMessageByProfile `json:"profile"`
	Content   string                                  `json:"content"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Message
		fmt.Println("[WriteMessage message] ", message)
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
		messageData := &Message{}
		err = json.Unmarshal(m, &messageData)
		if err != nil {
			logger.Log.Debug("error func ReadMessage,"+
				" method Unmarshal by path internal/entity/ws/client.go",
				zap.Error(err))
		}
		msg := &Message{
			ID:        c.ID,
			RoomID:    c.RoomID,
			UserID:    c.UserID,
			Type:      NoneMessage,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsDeleted: false,
			IsEdited:  false,
			Content:   messageData.Content,
		}
		h.Broadcast <- msg
	}
}
