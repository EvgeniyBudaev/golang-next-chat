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
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}
