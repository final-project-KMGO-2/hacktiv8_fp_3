package entity

import "time"

type Task struct {
	Title       string `json:"title"`      
	Description string `json:"description"`
	Deadline    time.Time `json:"deadline"`   
	Status      string `json:"status"`     
	Category    string `json:"category"`   
	UserID      int64  `json:"userId"`     
}
