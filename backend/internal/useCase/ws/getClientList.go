package ws

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/gofiber/fiber/v2"
)

type GetClientListUseCase struct {
	hub *ws.Hub
}

func NewGetClientListUseCase(h *ws.Hub) *GetClientListUseCase {
	return &GetClientListUseCase{
		hub: h,
	}
}

func (uc *GetClientListUseCase) GetClientList(ctx *fiber.Ctx) ([]*ws.ClientResponse, error) {
	clients := make([]*ws.ClientResponse, 0)
	roomId := ctx.Params("roomId")
	if _, ok := uc.hub.Rooms[roomId]; !ok {
		return clients, nil
	}
	for _, c := range uc.hub.Rooms[roomId].Clients {
		clients = append(clients, &ws.ClientResponse{
			ID:       c.ID,
			Username: c.Username,
		})
	}
	return clients, nil
}
