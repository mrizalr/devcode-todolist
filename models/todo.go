package models

import (
	"time"
)

type Todo struct {
	ID         uint      `json:"id" gorm:"primary key"`
	ActivityID uint      `json:"activity_group_id"`
	Title      string    `json:"title" gorm:"type:varchar(30)"`
	IsActive   bool      `json:"is_active"`
	Priority   string    `json:"priority" gorm:"type:varchar(30)"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Activity   Activity  `gorm:"foreignkey:ActivityID"`
}

// func (t *Todo) BeforeCreate(tx *gorm.DB) error {
// 	t.CreatedAt = time.Now()
// 	return nil
// }

// func (t *Todo) BeforeUpdate(tx *gorm.DB) error {
// 	t.UpdatedAt = time.Now()
// 	return nil
// }
