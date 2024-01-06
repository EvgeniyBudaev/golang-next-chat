package ws

import "github.com/gofiber/fiber/v2"

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(ctx *fiber.Ctx) {
	var req CreateRoomRequest
	if err := ctx.BodyParser(&req); err != nil {
		//return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Хранение информации по комнате в памяти, а не в БД
	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}
	ctx.Status(fiber.StatusCreated).JSON(req)
}
