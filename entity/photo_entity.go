package entity

type Photo struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint64 `gorm:"foreignKey" json:"user_id"`
	User     *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	BaseModel
}

type PhotoCreate struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserID   uint64 `gorm:"foreignKey" json:"user_id"`
}

type PhotoUpdate struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserID   uint64 `gorm:"foreignKey" json:"user_id"`
}
