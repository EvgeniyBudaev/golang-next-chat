package room

import (
	"context"
	"database/sql"
	errorEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/error"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/pagination"
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/searching"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/ws"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

type DBRoom interface {
	CreateRoom(ctx context.Context, r *ws.Room) (*ws.Room, error)
	InsertRoomProfiles(ctx context.Context, roomID int64, senderId int64, receiverId int64) error
	SelectRoomList(ctx context.Context, qp *ws.QueryParamsRoomList) ([]*ws.RoomWithProfileResponse, error)
	SelectRoomListByProfile(ctx context.Context, profileId int64) ([]*ws.RoomWithProfileResponse, error)
	AddRoomProfile(ctx context.Context, r *ws.RoomProfile) (*ws.RoomProfile, error)
	FindProfile(ctx context.Context, userId string) (*profileEntity.Profile, error)
	SelectUserList(ctx context.Context) ([]*ws.Client, error)
	AddMessage(ctx context.Context, m *ws.Message) (*ws.Message, error)
	SelectMessageList(ctx context.Context, roomId int64, page uint64, limit uint64) (*ws.ResponseMessageList, error)
	CheckIfCommonRoomExists(ctx context.Context, senderId int64, receiverId int64) (bool, int64, error)
}

type PGRoomDB struct {
	db *sql.DB
}

func NewPGRoomDB(db *sql.DB) *PGRoomDB {
	return &PGRoomDB{db: db}
}

func (pg *PGRoomDB) CreateRoom(ctx context.Context, r *ws.Room) (*ws.Room, error) {
	query := "INSERT INTO rooms (room_name, title) VALUES ($1, $2) RETURNING id"
	err := pg.db.QueryRowContext(ctx, query, r.RoomName, r.Title).Scan(&r.ID)
	if err != nil {
		logger.Log.Debug("error func CreateRoom, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	return r, nil
}

func (pg *PGRoomDB) InsertRoomProfiles(ctx context.Context, roomID int64, senderId int64, receiverId int64) error {
	query := "INSERT INTO rooms_profiles (room_id, profile_id) VALUES ($1, $2), ($1, $3)"
	_, err := pg.db.ExecContext(ctx, query, roomID, senderId, receiverId)
	if err != nil {
		logger.Log.Debug("error func InsertRoomProfiles, method ExecContext by path internal/db/room/room.go",
			zap.Error(err))
		return err
	}
	return nil
}

func (pg *PGRoomDB) SelectRoomList(
	ctx context.Context, qp *ws.QueryParamsRoomList) ([]*ws.RoomWithProfileResponse, error) {
	query := "SELECT id, room_name, title FROM rooms"
	query = searching.ApplySearch(query, "room_name", qp.Search) // search
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
		err := rows.Scan(&data.ID, &data.RoomName, &data.Title)
		if err != nil {
			logger.Log.Debug("error func SelectRoomList, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}
		list = append(list, &data)
	}
	return list, nil
}

func (pg *PGRoomDB) SelectRoomListByProfile(
	ctx context.Context, profileId int64) ([]*ws.RoomWithProfileResponse, error) {
	query := "SELECT rooms.id, rooms.room_name, rooms.title " +
		"FROM rooms " +
		"JOIN rooms_profiles ON rooms.id = rooms_profiles.room_id " +
		"JOIN profiles ON profiles.id = rooms_profiles.profile_id " +
		"WHERE profiles.id = $1"
	rows, err := pg.db.QueryContext(ctx, query, profileId)
	if err != nil {
		logger.Log.Debug(
			"error func SelectRoomListByProfile, method QueryContext by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*ws.RoomWithProfileResponse, 0)
	for rows.Next() {
		data := ws.RoomWithProfileResponse{}
		err := rows.Scan(&data.ID, &data.RoomName, &data.Title)
		if err != nil {
			logger.Log.Debug("error func SelectRoomListByProfile, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}
		list = append(list, &data)
	}
	return list, nil
}

func (pg *PGRoomDB) AddRoomProfile(ctx context.Context, r *ws.RoomProfile) (*ws.RoomProfile, error) {
	// Проверяем, существует ли запись
	checkQuery := "SELECT id FROM rooms_profiles WHERE room_id = $1 AND profile_id = $2"
	var existingID int
	err := pg.db.QueryRowContext(ctx, checkQuery, r.RoomID, r.ProfileID).Scan(&existingID)
	if err == nil {
		// Запись уже существует, возвращаем ошибку или какой-то код, чтобы обработать это
		logger.Log.Debug("error func AddRoomProfile, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		msg := errors.Wrap(err, "record already exists")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	// Запись не существует, выполняем операцию вставки
	insertQuery := "INSERT INTO rooms_profiles (room_id, profile_id) VALUES ($1, $2) RETURNING id"
	err = pg.db.QueryRowContext(ctx, insertQuery, r.RoomID, r.ProfileID).Scan(&r.ID)
	if err != nil {
		logger.Log.Debug("error func AddRoomProfile, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	return r, nil
}

func (pg *PGRoomDB) FindProfile(ctx context.Context, userId string) (*profileEntity.Profile, error) {
	p := profileEntity.Profile{}
	query := `SELECT id, user_id, username, first_name, last_name, email, created_at, updated_at, is_deleted,
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
	err := row.Scan(&p.ID, &p.UserID, &p.Username, &p.Firstname, &p.Lastname, &p.Email, &p.CreatedAt,
		&p.UpdatedAt, &p.IsDeleted, &p.IsEnabled)
	if err != nil {
		logger.Log.Debug("error func FindProfile, method Scan by path internal/db/room/room.go",
			zap.Error(err))
		return nil, err
	}
	return &p, nil
}

func (pg *PGRoomDB) SelectUserList(ctx context.Context) ([]*ws.Client, error) {
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

func (pg *PGRoomDB) AddMessage(ctx context.Context, m *ws.Message) (*ws.Message, error) {
	query := "INSERT INTO room_messages (room_id, user_id, type, created_at, updated_at, is_deleted, is_edited," +
		" is_joined, is_left, content) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	err := pg.db.QueryRowContext(ctx, query, m.RoomID, m.UserID, m.Type, m.CreatedAt, m.UpdatedAt,
		m.IsDeleted, m.IsEdited, m.IsJoined, m.IsLeft, m.Content).Scan(&m.ID)
	if err != nil {
		logger.Log.Debug("error func AddMessage, method QueryRowContext by path internal/db/room/room.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	return m, nil
}

func (pg *PGRoomDB) SelectMessageList(
	ctx context.Context, roomId int64, page uint64, limit uint64) (*ws.ResponseMessageList, error) {
	query := "SELECT rm.id, rm.room_id, rm.user_id, rm.type, rm.created_at, rm.updated_at, rm.is_deleted, " +
		"rm.is_edited, rm.is_joined, rm.is_left, rm.content, " +
		"p.id, p.first_name, p.last_name " +
		"FROM room_messages rm " +
		"JOIN profiles p ON rm.user_id = p.user_id " +
		"WHERE rm.room_id = $1 " +
		"ORDER BY rm.created_at DESC"
	countQuery := "SELECT COUNT(*) FROM room_messages"
	// get totalItems
	totalItems, err := pagination.GetTotalItems(ctx, pg.db, countQuery)
	if err != nil {
		logger.Log.Debug("error while counting SelectList. error in method GetTotalItems", zap.Error(err))
		return nil, err
	}
	// pagination
	query = pagination.ApplyPagination(query, page, limit)
	countQuery = pagination.ApplyPagination(countQuery, page, limit)
	// get list
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
		err := rows.Scan(&data.ID, &data.RoomID, &data.UserID, &data.Type, &data.CreatedAt, &data.UpdatedAt,
			&data.IsDeleted, &data.IsEdited, &data.IsJoined, &data.IsLeft, &data.Content,
			&profile.ID, &profile.Firstname, &profile.Lastname)
		if err != nil {
			logger.Log.Debug("error func SelectMessageList, method Scan by path internal/db/room/room.go",
				zap.Error(err))
			continue
		}
		data.Profile = &profile
		list = append(list, &data)
	}
	paging := pagination.GetPagination(limit, page, totalItems)
	response := ws.ResponseMessageList{
		Pagination: paging,
		Content:    list,
	}
	return &response, nil
}

// CheckIfCommonRoomExists - проверка на наличие общей комнаты
func (pg *PGRoomDB) CheckIfCommonRoomExists(
	ctx context.Context, senderId int64, receiverId int64) (bool, int64, error) {
	var roomID int64
	query := "SELECT room_id " +
		"FROM rooms_profiles WHERE profile_id = $1 INTERSECT SELECT room_id FROM rooms_profiles WHERE profile_id = $2"
	err := pg.db.QueryRowContext(ctx, query, senderId, receiverId).Scan(&roomID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, 0, nil
		}
		return false, 0, err
	}
	return true, roomID, nil
}
