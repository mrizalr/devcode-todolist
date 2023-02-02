package models

import (
	"time"
)

type Activity struct {
	ID        uint      `json:"id" gorm:"primary key"`
	Title     string    `json:"title" gorm:"type:varchar(30)"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func (a *Activity) BeforeCreate(tx *gorm.DB) error {
// 	a.CreatedAt = time.Now()
// 	return nil
// }

// func (a *Activity) BeforeUpdate(tx *gorm.DB) error {
// 	a.UpdatedAt = time.Now()
// 	return nil
// }
