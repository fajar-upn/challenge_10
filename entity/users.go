package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey; not null; type:uuid"`
	Username     string    `json:"username" gorm:"not null; type:text;"`
	Password     string    `json:"password" gorm:"not null; type:text;"`
	Access_level string    `json:"access_level" gorm:"not null; type:text;"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type PayloadUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
