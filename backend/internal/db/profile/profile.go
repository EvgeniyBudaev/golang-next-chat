package profile

import (
	"database/sql"
	errorEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/error"
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

type DBProfile interface {
	Create(cf *fiber.Ctx, p *profileEntity.Profile) (*profileEntity.Profile, error)
	FindByUsername(ctx *fiber.Ctx, username string) (*profileEntity.Profile, error)
	AddImage(cf *fiber.Ctx, p *profileEntity.ImageProfile) (*profileEntity.ImageProfile, error)
	SelectListImage(cf *fiber.Ctx, profileID int) ([]*profileEntity.ImageProfile, error)
}

type PGProfileDB struct {
	db *sql.DB
}

func NewPGProfileDB(db *sql.DB) *PGProfileDB {
	return &PGProfileDB{db: db}
}

func (pg *PGProfileDB) Create(cf *fiber.Ctx, p *profileEntity.Profile) (*profileEntity.Profile, error) {
	ctx := cf.Context()
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func Create, method Begin by path internal/db/profile/profile.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO profiles (uuid, user_id, username, first_name, last_name, email, created_at, updated_at," +
		" is_deleted, is_enabled) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	err = tx.QueryRowContext(ctx, query, p.UUID, p.UserID, p.Username, p.Firstname, p.Lastname, p.Email,
		p.CreatedAt, p.UpdatedAt, p.IsDeleted, p.IsEnabled).Scan(&p.ID)
	if err != nil {
		logger.Log.Debug("error func Create, method QueryRowContext by path internal/db/profile/profile.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}

	tx.Commit()
	return p, nil
}

func (pg *PGProfileDB) FindByUsername(ctx *fiber.Ctx, username string) (*profileEntity.Profile, error) {
	p := profileEntity.Profile{}
	query := `SELECT id, uuid, user_id, username, first_name, last_name, email, created_at, updated_at, is_deleted,
       is_enabled
			  FROM profiles
			  WHERE username = $1`
	row := pg.db.QueryRowContext(ctx.Context(), query, username)
	if row == nil {
		err := errors.New("no rows found")
		logger.Log.Debug(
			"error func FindByUsername, method QueryRowContext by path internal/db/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	err := row.Scan(&p.ID, &p.UUID, &p.UserID, &p.Username, &p.Firstname, &p.Lastname, &p.Email, &p.CreatedAt,
		&p.UpdatedAt, &p.IsDeleted, &p.IsEnabled)
	if err != nil {
		logger.Log.Debug("error func FindByUsername, method Scan by path internal/db/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	return &p, nil
}

func (pg *PGProfileDB) AddImage(cf *fiber.Ctx, p *profileEntity.ImageProfile) (*profileEntity.ImageProfile, error) {
	ctx := cf.Context()
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func AddImage, method Begin by path internal/db/profile/profile.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO profile_images (profile_id, uuid, name, url, size, created_at, updated_at, is_deleted," +
		" is_enabled) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	err = tx.QueryRowContext(ctx, query, p.ProfileID, p.UUID, p.Name, p.Url, p.Size, p.CreatedAt, p.UpdatedAt,
		p.IsDeleted, p.IsEnabled).Scan(&p.ID)
	if err != nil {
		logger.Log.Debug("error func AddImage, method QueryRowContext by path internal/db/profile/profile.go",
			zap.Error(err))
		msg := errors.Wrap(err, "bad request")
		err = errorEntity.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
	tx.Commit()
	return p, nil
}

func (pg *PGProfileDB) SelectListImage(cf *fiber.Ctx, profileID int) ([]*profileEntity.ImageProfile, error) {
	ctx := cf.Context()
	query := `SELECT id, profile_id, uuid, name, url, size, created_at, updated_at, is_deleted, is_enabled
	FROM profile_images
	WHERE profile_id = $1`
	rows, err := pg.db.QueryContext(ctx, query, profileID)
	if err != nil {
		logger.Log.Debug("error func SelectListImage, method QueryContext by path internal/db/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*profileEntity.ImageProfile, 0)
	for rows.Next() {
		data := profileEntity.ImageProfile{}
		err := rows.Scan(&data.ID, &data.ProfileID, &data.UUID, &data.Name, &data.Url, &data.Size,
			&data.CreatedAt, &data.UpdatedAt, &data.IsDeleted, &data.IsEnabled)
		if err != nil {
			logger.Log.Debug("error func SelectListImage, method Scan by path internal/db/profile/profile.go",
				zap.Error(err))
			continue
		}
		list = append(list, &data)
	}
	return list, nil
}
