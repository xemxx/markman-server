package model

import (
	"fmt"
	"markman-server/tools/config"
	"markman-server/tools/logs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Db .
var Db *gorm.DB

//Model .
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	dbCfg := config.Cfg.GetStringMapString("database")
	Db, err := gorm.Open(dbCfg["type"], fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		dbCfg["user"],
		dbCfg["password"],
		dbCfg["host"],
		dbCfg["database"],
		"Asia%2FShanghai"))
	if err != nil {
		logs.Log(err.Error())
		return
	}
	Db.SingularTable(true)
	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}

//CloseDB .
func CloseDB() {
	defer Db.Close()
}

//Test .
func Test() {
	user := &User{}
	Db.Select("id")
	Db.Where(&User{Username: "admin"})
	Db.First(user)
	fmt.Println(user)
}
