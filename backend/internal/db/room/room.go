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
	AddUser(c *ws.Client) (*ws.Client, error)
	SelectUserList() ([]*ws.Client, error)
	AddMessage(m *ws.Message) (*ws.Message, error)
	SelectMessageList(cf *fiber.Ctx, roomId int64) ([]*ws.Message, error)
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

func (pg *PGRoomDB) AddUser(c *ws.Client) (*ws.Client, error) {
	//ctx := cf.Context()
	ctx := context.Background()
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func AddUser, method Begin by path internal/db/room/room.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO room_users (room_id, user_id) VALUES ($1, $2) RETURNING id"
	err = tx.QueryRowContext(ctx, query, c.RoomID, c.UserID).Scan(&c.ID)
	if err != nil {
		logger.Log.Debug("error func AddUser, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	tx.Commit()
	return c, nil
}

func (pg *PGRoomDB) SelectUserList() ([]*ws.Client, error) {
	//ctx := cf.Context()
	ctx := context.Background()
	query := `SELECT id, room_id, user_id FROM room_users`
	rows, err := pg.db.QueryContext(ctx, query)
	if err != nil {
		logger.Log.Debug("error func SelectUserList, method QueryContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*ws.Client, 0)
	for rows.Next() {
		data := ws.Client{}
		err := rows.Scan(&data.ID, &data.RoomID, &data.UserID)
		if err != nil {
			logger.Log.Debug("error func SelectUserList, method Scan by path internal/db/room/room.go",
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
	query := "INSERT INTO room_messages (room_id, user_id, content) VALUES ($1, $2, $3) RETURNING id"
	err = tx.QueryRowContext(ctx, query, m.RoomID, m.UserID, m.Content).Scan(&m.ID)
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

func (pg *PGRoomDB) SelectMessageList(cf *fiber.Ctx, roomId int64) ([]*ws.Message, error) {
	ctx := cf.Context()
	query := `SELECT id, room_id, user_id, content FROM room_messages WHERE room_id = $1`
	rows, err := pg.db.QueryContext(ctx, query, roomId)
	if err != nil {
		logger.Log.Debug("error func SelectMessageList, method QueryContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*ws.Message, 0)
	for rows.Next() {
		data := ws.Message{}
		err := rows.Scan(&data.ID, &data.RoomID, &data.UserID, &data.Content)
		if err != nil {
			logger.Log.Debug("error func SelectMessageList, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}
		list = append(list, &data)
	}
	return list, nil
}
