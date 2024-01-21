package profile

import (
	"database/sql"
	errorEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/error"
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/searching"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

type DBProfile interface {
	Create(ctf *fiber.Ctx, p *profileEntity.Profile) (*profileEntity.Profile, error)
	FindByUsername(ctf *fiber.Ctx, username string) (*profileEntity.Profile, error)
	SelectProfileList(ctf *fiber.Ctx, qp *profileEntity.QueryParamsProfileList) ([]*profileEntity.Profile, error)
	AddImage(ctf *fiber.Ctx, p *profileEntity.ImageProfile) (*profileEntity.ImageProfile, error)
	SelectListImage(ctf *fiber.Ctx, profileID int) ([]*profileEntity.ImageProfile, error)
}

type PGProfileDB struct {
	db *sql.DB
}

func NewPGProfileDB(db *sql.DB) *PGProfileDB {
	return &PGProfileDB{db: db}
}

func (pg *PGProfileDB) Create(ctf *fiber.Ctx, p *profileEntity.Profile) (*profileEntity.Profile, error) {
	ctx := ctf.Context()
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func Create, method Begin by path internal/db/profile/profile.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO profiles (user_id, username, first_name, last_name, email, created_at, updated_at," +
		" is_deleted, is_enabled) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	err = tx.QueryRowContext(ctx, query, p.UserID, p.Username, p.Firstname, p.Lastname, p.Email,
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

func (pg *PGProfileDB) FindByUsername(ctf *fiber.Ctx, username string) (*profileEntity.Profile, error) {
	p := profileEntity.Profile{}
	query := `SELECT id, user_id, username, first_name, last_name, email, created_at, updated_at, is_deleted,
       is_enabled
			  FROM profiles
			  WHERE username = $1`
	row := pg.db.QueryRowContext(ctf.Context(), query, username)
	if row == nil {
		err := errors.New("no rows found")
		logger.Log.Debug(
			"error func FindByUsername, method QueryRowContext by path internal/db/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	err := row.Scan(&p.ID, &p.UserID, &p.Username, &p.Firstname, &p.Lastname, &p.Email, &p.CreatedAt,
		&p.UpdatedAt, &p.IsDeleted, &p.IsEnabled)
	if err != nil {
		logger.Log.Debug("error func FindByUsername, method Scan by path internal/db/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	return &p, nil
}

func (pg *PGProfileDB) SelectProfileList(ctf *fiber.Ctx, qp *profileEntity.QueryParamsProfileList) ([]*profileEntity.Profile, error) {
	query := "SELECT id, user_id, username, first_name, last_name, email, created_at, updated_at, is_deleted," +
		" is_enabled FROM profiles"
	query = searching.ApplySearch(query, "username", qp.Search) // search
	rows, err := pg.db.QueryContext(ctf.Context(), query)
	if err != nil {
		logger.Log.Debug(
			"error func SelectProfileList, method QueryContext by path internal/db/profile/profile.go",
			zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*profileEntity.Profile, 0)
	for rows.Next() {
		data := profileEntity.Profile{}
		err := rows.Scan(&data.ID, &data.UserID, &data.Username, &data.Firstname, &data.Lastname, &data.Email,
			&data.CreatedAt, &data.UpdatedAt, &data.IsDeleted, &data.IsEnabled)
		if err != nil {
			logger.Log.Debug("error func SelectProfileList, method Scan by path internal/db/profile/profile.go",
				zap.Error(err))
			continue
		}
		list = append(list, &data)
	}
	return list, nil
}

func (pg *PGProfileDB) AddImage(ctf *fiber.Ctx, p *profileEntity.ImageProfile) (*profileEntity.ImageProfile, error) {
	ctx := ctf.Context()
	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Debug("error func AddImage, method Begin by path internal/db/profile/profile.go", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "INSERT INTO profile_images (profile_id, name, url, size, created_at, updated_at, is_deleted," +
		" is_enabled) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err = tx.QueryRowContext(ctx, query, p.ProfileID, p.Name, p.Url, p.Size, p.CreatedAt, p.UpdatedAt,
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

func (pg *PGProfileDB) SelectListImage(ctf *fiber.Ctx, profileID int) ([]*profileEntity.ImageProfile, error) {
	ctx := ctf.Context()
	query := `SELECT id, profile_id, name, url, size, created_at, updated_at, is_deleted, is_enabled
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
		err := rows.Scan(&data.ID, &data.ProfileID, &data.Name, &data.Url, &data.Size,
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
