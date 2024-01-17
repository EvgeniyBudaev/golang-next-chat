package room

import (
	"context"
	"database/sql"
	errorEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/error"
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

type DBRoom interface {
	CreateRoom(cf *fiber.Ctx, p *ws.Room) (*ws.Room, error)
	SelectRoomList(cf *fiber.Ctx) ([]*ws.RoomWithProfileResponse, error)
	AddRoomProfile(roomID int64, profileID int64) (*ws.RoomProfile, error)
	FindProfile(userId string) (*profileEntity.Profile, error)
	SelectUserList() ([]*ws.Client, error)
	AddMessage(m *ws.Message) (*ws.Message, error)
	SelectMessageList(cf *fiber.Ctx, roomId int64) ([]*ws.ResponseMessage, error)
}

type PGRoomDB struct {
	db *sql.DB
}

func NewPGRoomDB(db *sql.DB) *PGRoomDB {
	return &PGRoomDB{db: db}
}

func (pg *PGRoomDB) CreateRoom(cf *fiber.Ctx, r *ws.Room) (*ws.Room, error) {
	ctx := cf.Context()
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func Create, method Begin by path internal/db/room/room.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO rooms (uuid) VALUES ($1) RETURNING id"
	err = tx.QueryRowContext(ctx, query, r.UUID).Scan(&r.ID)
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

func (pg *PGRoomDB) SelectRoomList(cf *fiber.Ctx) ([]*ws.RoomWithProfileResponse, error) {
	ctx := cf.Context()
	query := "SELECT rooms.id, rooms.uuid, profiles.uuid, profiles.first_name, profiles.last_name " +
		"FROM rooms " +
		"JOIN rooms_profiles " +
		"ON rooms.id = rooms_profiles.room_id JOIN profiles ON profiles.id = rooms_profiles.profile_id"
	rows, err := pg.db.QueryContext(ctx, query)
	if err != nil {
		logger.Log.Debug("error func SelectRoomList, method QueryContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*ws.RoomWithProfileResponse, 0)
	for rows.Next() {
		data := ws.RoomWithProfileResponse{}
		profile := profileEntity.ResponseProfileForRoom{}
		err := rows.Scan(&data.ID, &data.UUID, &profile.UUID, &profile.Firstname, &profile.Lastname)
		if err != nil {
			logger.Log.Debug("error func SelectRoomList, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}
		data.Profile = &profile
		list = append(list, &data)
	}
	return list, nil
}

func (pg *PGRoomDB) AddRoomProfile(roomID int64, profileID int64) (*ws.RoomProfile, error) {
	//ctx := cf.Context()
	ctx := context.Background()
	r := ws.RoomProfile{}
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func AddRoomProfile, method Begin by path internal/db/room/room.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO rooms_profiles (room_id, profile_id) VALUES ($1, $2) RETURNING id"
	err = tx.QueryRowContext(ctx, query, roomID, profileID).Scan(&r.ID)
	if err != nil {
		logger.Log.Debug("error func AddRoomProfile, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	tx.Commit()
	return &r, nil
}

func (pg *PGRoomDB) FindProfile(userId string) (*profileEntity.Profile, error) {
	//ctx := cf.Context()
	ctx := context.Background()
	p := profileEntity.Profile{}
	query := `SELECT id, uuid, user_id, username, first_name, last_name, email, created_at, updated_at, is_deleted,
       is_enabled
			  FROM profiles
			  WHERE user_id = $1`
	row := pg.db.QueryRowContext(ctx, query, userId)
	if row == nil {
		err := errors.New("no rows found")
		logger.Log.Debug("error func FindProfile, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	err := row.Scan(&p.ID, &p.UUID, &p.UserID, &p.Username, &p.Firstname, &p.Lastname, &p.Email, &p.CreatedAt,
		&p.UpdatedAt, &p.IsDeleted, &p.IsEnabled)
	if err != nil {
		logger.Log.Debug("error func FindProfile, method Scan by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return &p, nil
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
	query := "INSERT INTO room_messages (uuid, room_id, user_id, type, created_at, updated_at, is_deleted, is_edited," +
		" content) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	err = tx.QueryRowContext(ctx, query, m.UUID, m.RoomID, m.UserID, m.Type, m.CreatedAt, m.UpdatedAt, m.IsDeleted,
		m.IsEdited, m.Content).Scan(&m.ID)
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

func (pg *PGRoomDB) SelectMessageList(cf *fiber.Ctx, roomId int64) ([]*ws.ResponseMessage, error) {
	ctx := cf.Context()
	query := "SELECT rm.uuid, rm.room_id, rm.user_id, rm.type, rm.created_at, rm.updated_at, rm.is_deleted, " +
		"rm.is_edited, rm.content, " +
		"p.uuid, p.first_name, p.last_name " +
		"FROM room_messages rm " +
		"JOIN profiles p ON rm.user_id = p.user_id " +
		"WHERE rm.room_id = $1 " +
		"ORDER BY rm.created_at ASC"
	rows, err := pg.db.QueryContext(ctx, query, roomId)
	if err != nil {
		logger.Log.Debug("error func SelectMessageList, method QueryContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*ws.ResponseMessage, 0)
	for rows.Next() {
		data := ws.ResponseMessage{}
		profile := profileEntity.ResponseMessageByProfile{}
		err := rows.Scan(&data.UUID, &data.RoomID, &data.UserID, &data.Type, &data.CreatedAt, &data.UpdatedAt,
			&data.IsDeleted, &data.IsEdited, &data.Content,
			&profile.UUID, &profile.Firstname, &profile.Lastname)
		if err != nil {
			logger.Log.Debug("error func SelectMessageList, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}
		data.Profile = &profile
		list = append(list, &data)
	}
	return list, nil
}
