package profile

import (
	"time"
)

type Profile struct {
	ID        int64           `json:"id"`
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
	ID        int64  `json:"id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

type ImageProfile struct {
	ID        int64     `json:"id"`
	ProfileID int64     `json:"profileId"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDeleted bool      `json:"isDeleted"`
	IsEnabled bool      `json:"isEnabled"`
}
