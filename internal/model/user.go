package model

import "time"

type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"not null" json:"name"`
	Email         string    `gorm:"uniqueIndex;not null" json:"email"`
	Password      string    `json:"-"`
	Role          Role      `gorm:"type:text;check:role IN ('user','admin');default:'user'"`
	EmailVerified bool      `gorm:"default:false" json:"is_email_verified"`
	VerifyToken   string    `json:"-"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
