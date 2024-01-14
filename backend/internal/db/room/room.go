package room

import (
	"context"
	"database/sql"
	errorEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/error"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

type DBRoom interface {
	Create(cf *fiber.Ctx, p *ws.Room) (*ws.Room, error)
	SelectRoomList(cf *fiber.Ctx) ([]*ws.Room, error)
	AddClient(c *ws.Client) (*ws.Client, error)
	SelectClientList() ([]*ws.Client, error)
	AddMessage(m *ws.Message) (*ws.Message, error)
}

type PGRoomDB struct {
	db *sql.DB
}

func NewPGRoomDB(db *sql.DB) *PGRoomDB {
	return &PGRoomDB{db: db}
}

func (pg *PGRoomDB) Create(cf *fiber.Ctx, r *ws.Room) (*ws.Room, error) {
	ctx := cf.Context()
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func Create, method Begin by path internal/db/room/room.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO rooms (name) VALUES ($1) RETURNING id"
	err = tx.QueryRowContext(ctx, query, r.Name).Scan(&r.ID)
	if err != nil {
		logger.Log.Debug("error func Create, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	tx.Commit()
	return r, nil
}

func (pg *PGRoomDB) SelectRoomList(cf *fiber.Ctx) ([]*ws.Room, error) {
	ctx := cf.Context()
	query := `SELECT id, name FROM rooms`
	rows, err := pg.db.QueryContext(ctx, query)
	if err != nil {
		logger.Log.Debug("error func SelectList, method QueryContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*ws.Room, 0)
	for rows.Next() {
		data := ws.Room{}
		err := rows.Scan(&data.ID, &data.Name)
		if err != nil {
			logger.Log.Debug("error func SelectList, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}
		list = append(list, &data)
	}
	return list, nil
}

func (pg *PGRoomDB) AddClient(c *ws.Client) (*ws.Client, error) {
	//ctx := cf.Context()
	ctx := context.Background()
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func AddClient, method Begin by path internal/db/room/room.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO room_clients (room_id, user_id, username) VALUES ($1, $2, $3) RETURNING id"
	err = tx.QueryRowContext(ctx, query, c.RoomID, c.UserID, c.Username).Scan(&c.ID)
	if err != nil {
		logger.Log.Debug("error func AddClient, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	tx.Commit()
	return c, nil
}

func (pg *PGRoomDB) SelectClientList() ([]*ws.Client, error) {
	//ctx := cf.Context()
	ctx := context.Background()
	query := `SELECT id, room_id, user_id, username FROM room_clients`
	rows, err := pg.db.QueryContext(ctx, query)
	if err != nil {
		logger.Log.Debug("error func SelectClientList, method QueryContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*ws.Client, 0)
	for rows.Next() {
		data := ws.Client{}
		err := rows.Scan(&data.ID, &data.RoomID, &data.UserID, &data.Username)
		if err != nil {
			logger.Log.Debug("error func SelectClientList, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}
		list = append(list, &data)
	}
	return list, nil
}

func (pg *PGRoomDB) AddMessage(m *ws.Message) (*ws.Message, error) {
	//ctx := cf.Context()
	ctx := context.Background()
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func AddMessage, method Begin by path internal/db/room/room.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO room_client_messages (room_id, client_id, content) VALUES ($1, $2, $3) RETURNING id"
	err = tx.QueryRowContext(ctx, query, m.RoomID, m.ClientID, m.Content).Scan(&m.ID)
	if err != nil {
		logger.Log.Debug("error func AddMessage, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	tx.Commit()
	return m, nil
}
