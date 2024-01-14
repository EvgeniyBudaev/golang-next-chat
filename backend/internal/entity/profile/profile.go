package profile

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID        int64           `json:"id"`
	UUID      uuid.UUID       `json:"uuid"`
	UserID    string          `json:"userId"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Images    []*ImageProfile `json:"images"`
}

type ResponseProfile struct {
	UUID      uuid.UUID       `json:"uuid"`
	UserID    string          `json:"userId"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Images    []*ImageProfile `json:"images"`
}

type ImageProfile struct {
	ID        int64     `json:"id"`
	ProfileID int64     `json:"profileId"`
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDeleted bool      `json:"isDeleted"`
	IsEnabled bool      `json:"isEnabled"`
}
