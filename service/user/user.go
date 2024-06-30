package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"

	"markman-server/model"
	"markman-server/tools/common"
)

// ExistUser .
func ExistUser(username, password string) (int, bool) {
	user := model.User{
		Username: username,
	}
	d := model.I().Where(&user, "username").First(&user)
	fmt.Println(user)
	if user.ID == 0 || d.Error != nil {
		fmt.Println(d.Error.Error())
		return 0, false
	}
	if !common.CheckPassword(user.Password, password) {
		return 0, false
	}
	return user.ID, true
}

// ExistUserByName .
func ExistUserByName(username string) bool {
	var user = model.User{Username: username}
	d := model.I().Where(&user).First(&user)
	if user.ID == 0 || d.Error != nil {
		return false
	}
	return true
}

// AddUser .
func AddUser(username, password string) (string, bool) {
	hash, err := common.NewPassword(password)
	if err != nil {
		slog.Info("generate password failed", "err", err)
		return "", false
	}
	user := model.User{
		Username: username,
		Password: hash,
		SC:       0,
		UUID:     uuid.New().String(),
	}
	if rows := model.I().Create(&user).RowsAffected; rows == 0 {
		slog.Info("create user failed", "err", model.I().Error.Error())
		return "", false
	}
	return user.UUID, true
}

// GetByPass .
func GetByPass(username, password string) (*model.User, bool) {
	user := model.User{
		Username: username,
	}
	d := model.I().Where(&user, "username").First(&user)
	fmt.Println(user)
	if user.ID == 0 || d.Error != nil {
		fmt.Println(d.Error.Error())
		return &user, false
	}
	if !common.CheckPassword(user.Password, password) {
		return &user, false
	}
	return &user, true
}

func UpdateSC(id, SC int) {
	model.I().Table("user").Where("id=?", id).Updates(map[string]interface{}{"SC": SC})
}

func Get(id int) model.User {
	var user model.User
	model.I().Where(`id=?`, id).First(&user)
	return user
}

// SaveToken .
func SaveToken(uid int, token string) bool {
	if rows := model.I().Model(&model.User{ID: uid}).Update("token", token).RowsAffected; rows == 0 {
		return false
	}
	return true
}

func GetLastSC(uid int) (int, error) {
	var user = model.User{ID: uid}
	user.GetByID()
	if user.SC > -1 {
		return user.SC, nil
	}
	return 0, errors.New("user not find")
}
