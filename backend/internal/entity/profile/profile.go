package profile

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID        int64           `json:"id"`
	UUID      uuid.UUID       `json:"uuid"`
	UserID    string          `json:"userId"`
	Username  string          `json:"username"`
	Firstname string          `json:"firstName"`
	Lastname  string          `json:"lastName"`
	Email     string          `json:"email"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	IsDeleted bool            `json:"isDeleted"`
	IsEnabled bool            `json:"isEnabled"`
	Images    []*ImageProfile `json:"images"`
}

type ResponseProfile struct {
	ID        int64           `json:"id"`
	UUID      uuid.UUID       `json:"uuid"`
	UserID    string          `json:"userId"`
	Username  string          `json:"username"`
	Firstname string          `json:"firstName"`
	Lastname  string          `json:"lastName"`
	Email     string          `json:"email"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	IsDeleted bool            `json:"isDeleted"`
	IsEnabled bool            `json:"isEnabled"`
	Images    []*ImageProfile `json:"images"`
}

type ResponseMessageByProfile struct {
	UUID      uuid.UUID `json:"uuid"`
	Firstname string    `json:"firstName"`
	Lastname  string    `json:"lastName"`
}

type ResponseProfileForRoom struct {
	UUID      uuid.UUID `json:"uuid"`
	Firstname string    `json:"firstName"`
	Lastname  string    `json:"lastName"`
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
