package model

import "time"

// User .
type User struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Username   string    `gorm:"unique_index;not null" json:"username"`
	Password   string    `json:"password"`
	Token      string    `json:"token"`
	CreateTime time.Time `json:"create_time"`
}

// //CheckUser .
// func CheckUser(username, password string) (int, bool) {
// 	var user User
// 	db.Select("id").Where(User{Username: username, Password: password}).First(&user)
// 	if user.ID > 0 {
// 		return user.ID, true
// 	}
// 	return 0, false
// }

// GetByID .
func (u *User) GetByID() {
	Db.Select(u).First(u)
}
