package ws

type Room struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RoomResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type QueryParamsJoinRoom struct {
	ClientID string `json:"clientId"`
	Username string `json:"username"`
}

type Hub struct {
	Clients    map[int64][]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int64][]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}
