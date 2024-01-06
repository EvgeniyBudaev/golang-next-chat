package ws

import (
	"bufio"
	"bytes"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
)

type JoinRoomUseCase struct {
	hub *ws.Hub
}

func NewJoinRoomUseCase(h *ws.Hub) *JoinRoomUseCase {
	return &JoinRoomUseCase{
		hub: h,
	}
}

type ResponseAdapter struct {
	response *fasthttp.Response
}

func (r *ResponseAdapter) Header() http.Header {
	return r.Header()
}

func (r *ResponseAdapter) Write(p []byte) (int, error) {
	return r.Write(p)
}

func (r *ResponseAdapter) WriteHeader(statusCode int) {
	r.response.SetStatusCode(statusCode)
}

func ConvertToHTTPRequest(fastHTTPRequest *fasthttp.Request) (*http.Request, error) {
	// Создаем буфер для записи запроса
	var buf bytes.Buffer
	// Создаем bufio.Writer, использующий буфер
	writer := bufio.NewWriter(&buf)
	// Пишем данные из fasthttp.Request в bufio.Writer
	err := fastHTTPRequest.Write(writer)
	if err != nil {
		return nil, err
	}
	// Принудительно сбрасываем оставшиеся данные из буфера в bufio.Writer
	writer.Flush()
	// Создаем http.Request, используя данные из буфера
	httpRequest, err := http.ReadRequest(bufio.NewReader(&buf))
	if err != nil {
		return nil, err
	}
	return httpRequest, nil
}

func (uc *JoinRoomUseCase) JoinRoom(ctx *fiber.Ctx) (string, error) {
	upgrader := websocket.Upgrader{
		// ReadBufferSize - размер буфера чтения
		ReadBufferSize: 1024,
		// WriteBufferSize - размер буфера записи
		WriteBufferSize: 1024,
		// функция проверки источника, которая проверяет запрос и возвращает логическое значение
		CheckOrigin: func(r *http.Request) bool {
			// TODO: убрать return true, после тестирования в postman
			//origin := r.Header.Get("Origin")
			//return origin == "http://localhost:3000"
			return true
		},
	}
	// адаптер для fasthttp.Response
	adapterResponse := &ResponseAdapter{response: ctx.Response()}
	// Преобразуем fasthttp.Request в http.Request
	httpRequest, err := ConvertToHTTPRequest(ctx.Request())
	conn, err := upgrader.Upgrade(adapterResponse, httpRequest, nil)
	if err != nil {
		return "", err
	}
	roomId := ctx.Params("roomId")
	params := ws.QueryParamsJoinRoom{}
	if err := ctx.QueryParser(&params); err != nil {
		logger.Log.Debug("error in method ctx.QueryParser", zap.Error(err))
		return "", err
	}
	cl := &ws.Client{
		ID:       params.ClientID,
		RoomID:   roomId,
		Username: params.Username,
		Conn:     conn,
		Message:  make(chan *ws.Message),
	}
	m := &ws.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomId,
		Username: params.Username,
	}
	// Register a new client through the register channel
	uc.hub.Register <- cl
	// Broadcast that message
	uc.hub.Broadcast <- m
	go cl.WriteMessage()
	cl.ReadMessage(uc.hub)
	return "is joined", nil
}
