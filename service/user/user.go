package user

import (
	"errors"
	"markman-server/tools/logs"

	"markman-server/model"
	"markman-server/tools/common"
	"time"
)

var db = model.Db

//ExistUser .
func ExistUser(username, password string) (int, bool) {
	user := model.User{
		Username: username,
	}
	db.Where(user).First(&user)
	if user.ID == 0 || db.Error != nil || !common.CheckPassword(user.Password, password) {
		return 0, false
	}
	return user.ID, true
}

// ExistUserByName .
func ExistUserByName(username string) bool {
	var user = model.User{Username: username}
	db.Where(user).First(&user)
	if user.ID == 0 || db.Error != nil {
		return false
	}
	return true
}

// AddUser .
func AddUser(username, password string) bool {
	hash, err := common.NewPassword(password)
	if err != nil {
		logs.Log("generate password failed: err: "+ err.Error())
		return false
	}
	user := model.User{
		Username:   username,
		Password:   hash,
		CreateTime: time.Now(),
	}
	if rows := db.Create(&user).RowsAffected; rows == 0 {
		logs.Log(db.Error.Error())
		return false
	}
	return true
}

// GetUserInfo .
func GetUserInfo(uid int) (model.User, error) {
	var user model.User
	db.Select("username,create_time").Where(&model.User{
		ID: uid,
	}).First(&user)
	if user.Username != "" {
		return user, nil
	}
	return model.User{}, errors.New("user not find")
}

// SaveToken .
func SaveToken(uid int, token string) bool {
	if rows := db.Model(&model.User{ID: uid}).Update("token", token).RowsAffected; rows == 0 {
		return false
	}
	return true
}
