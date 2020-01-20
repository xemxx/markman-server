package model

import (
	"time"
)

// User .
type User struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username   string    `gorm:"unique;not null" json:"username"`
	Password   string    `json:"password"`
	Token      string    `json:"token"`
	CreateTime time.Time `json:"create_time"`
}

// GetByID .
func (u *User) GetByID() {
	Db.Where(u).First(u)
}
