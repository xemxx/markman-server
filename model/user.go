package model

import "time"

// User .
type User struct {
	ID        int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username  string `gorm:"uniqueIndex;not null;size:255" json:"username"`
	UUID      string `gorm:"not null;size:32" json:"uuid"`
	Password  string `gorm:"not null" json:"password"`
	Token     string `gorm:"size:255" json:"token"`
	SC        int    `gorm:"column:SC" json:"SC"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetByID .
func (u *User) GetByID() {
	Db.Where(u).First(u)
}

func (u *User) TableName() string {
	return "user"
}
