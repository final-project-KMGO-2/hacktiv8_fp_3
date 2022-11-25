package entity

type Category struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Type   string `json:"type"`
	UserID uint64 `json:"user_id"`
	BaseModel
}

type CategoryCreate struct {
	UserID uint64 `json:"user_id"`
	Type   string `json:"type"`
}

type CategoryPatch struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	UserID uint64 `json:"user_id"`
	Type   string `json:"type"`
}
