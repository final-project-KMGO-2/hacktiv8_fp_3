package entity

type SocialMedia struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         uint64 `gorm:"foreignKey" json:"user_id"`
	User           *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	BaseModel
}

type SocialMediaCreate struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
	UserID         uint64
}

type SocialMediaUpdate struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
	UserID         uint64 `json:"user_id"`
}
