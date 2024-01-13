package profile

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/db/profile"
	profileEntity "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
)

type CreateProfileRequest struct {
	Username string `json:"username"`
}

type UseCaseProfile struct {
	db *profile.PGProfileDB
}

func NewUseCaseProfile(db *profile.PGProfileDB) *UseCaseProfile {
	return &UseCaseProfile{db: db}
}

func (uc *UseCaseProfile) Create() (*profileEntity.Profile, error) {
	return uc.Create()
}
