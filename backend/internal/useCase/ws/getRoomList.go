package ws

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/gofiber/fiber/v2"
)

type GetRoomListUseCase struct {
	hub *ws.Hub
}

func NewGetRoomListUseCase(h *ws.Hub) *GetRoomListUseCase {
	return &GetRoomListUseCase{
		hub: h,
	}
}

func (uc *GetRoomListUseCase) GetRoomList(ctx *fiber.Ctx) ([]*ws.RoomResponse, error) {
	rooms := make([]*ws.RoomResponse, 0)
	for _, r := range uc.hub.Rooms {
		rooms = append(rooms, &ws.RoomResponse{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	return rooms, nil
}
