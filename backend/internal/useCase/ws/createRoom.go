package ws

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/gofiber/fiber/v2"
)

type CreateRoomUseCase struct {
	hub *ws.Hub
}

func NewCreateRoomUseCase(h *ws.Hub) *CreateRoomUseCase {
	return &CreateRoomUseCase{
		hub: h,
	}
}

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (uc *CreateRoomUseCase) CreateRoom(ctx *fiber.Ctx, r CreateRoomRequest) (string, error) {
	// Хранение информации по комнате в памяти, а не в БД
	uc.hub.Rooms[r.ID] = &ws.Room{
		ID:      r.ID,
		Name:    r.Name,
		Clients: make(map[string]*ws.Client),
	}
	//ctx.Status(fiber.StatusCreated).JSON(r)
	return "is created", nil
}
