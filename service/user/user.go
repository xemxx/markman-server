package user

import (
	"markman-server/model"
	"time"
)

var db = model.Db

//ExistUser .
func ExistUser(username, password string) (int, bool) {
	var user model.User
	db.First(&user)
	if user.ID > 0 {
		return user.ID, true
	}
	return 0, false
}

// ExistUserByName .
func ExistUserByName(username string) bool {
	return db.NewRecord(&model.User{Username: username})
	var user model.User
	model.Db.Where("username=?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

// AddUser .
func AddUser(username, password string) {
	db.Create(&model.User{
		Username:   username,
		Password:   password,
		CreateTime: time.Now(),
	})
}

// GetUserInfo .
func GetUserInfo(uid int) model.User {
	user := model.User{
		ID: uid,
	}
	db.Select("*").Where(user).First(&user)
	return user
}

// SaveToken .
func SaveToken(uid int) {
	db.Save(&model.User{ID: uid})
}
