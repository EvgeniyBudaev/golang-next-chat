package ws

import (
	"encoding/json"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/pagination"
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
	Page     uint64 `json:"page"`
	Limit    uint64 `json:"limit"`
	Conn     *websocket.Conn
	Content  chan *Content
}

type ClientResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type Content struct {
	Message           *Message             `json:"message"`
	MessageListByRoom *ResponseMessageList `json:"messageListByRoom"`
	Page              uint64               `json:"page"`
	Limit             uint64               `json:"limit"`
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
	IsJoined  bool        `json:"isJoined"`
	IsLeft    bool        `json:"isLeft"`
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
	IsJoined  bool                                    `json:"isJoined"`
	IsLeft    bool                                    `json:"isLeft"`
	Profile   *profileEntity.ResponseMessageByProfile `json:"profile"`
	Content   string                                  `json:"content"`
}

type ReadContent struct {
	Content string `json:"content"`
	Page    uint64 `json:"page"`
	Limit   uint64 `json:"limit"`
}

type ResponseMessageList struct {
	*pagination.Pagination
	Content []*ResponseMessage `json:"content"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		content, ok := <-c.Content
		//fmt.Println("[WriteMessage content] ", content)
		if !ok {
			return
		}
		c.Conn.WriteJSON(content)
	}
}

func (c *Client) ReadMessage(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, m, err := c.Conn.ReadMessage()
		//fmt.Println("[ReadMessage m] ", string(m))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		contentData := &ReadContent{}
		err = json.Unmarshal(m, &contentData)
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
			Content:   contentData.Content,
		}
		h.Broadcast <- &Content{
			Message: msg,
			Page:    contentData.Page,
			Limit:   contentData.Limit,
		}
	}
}
