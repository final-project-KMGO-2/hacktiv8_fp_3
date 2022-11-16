package entity

import (
	"time"

	"gorm.io/gorm"
)

// Base model that includes uint64 ID and created, updated, deleted timestamps
type BaseModel struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
