package entity

type Category struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	Type string `json:"type"`
	BaseModel
}

type CategoryCreate struct {
	Type string `json:"type" binding:"required"`
}

type CategoryPatch struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	Type string `json:"type" binding:"required"`
}
