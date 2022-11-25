package entity

import (
	"time"
)

// Base model that includes uint64 ID and created, updated, deleted timestamps
type BaseModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at"`
}
