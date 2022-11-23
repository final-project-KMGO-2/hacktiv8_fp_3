package entity

import "time"

type Task struct {
	Id          int       `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      bool      `json:"status"`
	CategoryID  string    `json:"category_id"`
	UserID      int64     `json:"user_id"`
	BaseModel
}
