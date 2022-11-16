package entity

type Comment struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	UserID  uint64 `json:"user_id"`
	User    *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	PhotoID uint64 `json:"photo_id"`
	Photo   *Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo,omitempty"`
	Message string `json:"message"`
	BaseModel
}

type CommentCreate struct {
	UserID  uint64 `json:"user_id"`
	PhotoID uint64 `json:"photo_id"`
	Message string `json:"message" binding:"required"`
}

type CommentUpdate struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	UserID  uint64 `json:"user_id"`
	Message string `json:"message"`
}
