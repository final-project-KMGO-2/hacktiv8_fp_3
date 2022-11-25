package entity

type Task struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      bool   `json:"status" binding:"boolean"`
	CategoryID  int    `json:"category_id" gorm:"foreignKey"`
	UserID      int    `json:"user_id"`
	User        User   `json:"user"`
	BaseModel
}

type TaskDetail struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      bool   `json:"status" binding:"boolean"`
	CategoryID  int    `json:"category_id"`
	UserID      int    `json:"user_id"`
	BaseModel
	User UserUpdate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}

type TaskCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id"`
	UserID      int    `json:"user_id"`
}

type TaskUpdate struct { // memakai task id di params
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type TaskStatusModifier struct {
	Status bool `json:"status" binding:"boolean"`
}

type TaskCategoryModifier struct {
	CategoryID int `json:"category_id"`
}
