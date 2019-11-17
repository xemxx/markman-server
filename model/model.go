package model

import (
	"fmt"
	"log"
	"markman-server/tools/config"

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

	db, err := gorm.Open(dbCfg["type"], fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		dbCfg["user"],
		dbCfg["password"],
		dbCfg["host"],
		dbCfg["database"],
		"Asia%2FShanghai"))

	if err != nil {
		log.Println(err)
	}

	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return tablePrefix + defaultTableName
	// }

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	Db = db
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
