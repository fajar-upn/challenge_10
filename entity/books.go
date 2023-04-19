package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Book struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey; not null; type:uuid"`
	User_id    uuid.UUID `json:"user_id" gorm:"not null; type:uuid"`
	Book_name  string    `json:"book_name" gorm:"not null; type:text;"`
	Author     string    `json:"author" gorm:"not null; type:varchar(150)"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type PayloadBook struct {
	User_id   uuid.UUID `json:"user_id"`
	Book_name string    `json:"book_name"`
	Author    string    `json:"author"`
}

type UpdateBoook struct {
	ID        string    `json:"id"`
	Book_name string    `json:"book_name"`
	Author    string    `json:"author"`
	Update_at time.Time `json:"updated_at"`
}
