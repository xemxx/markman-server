package models

import (
	"fmt"
	"log"
	"markman-server/tools/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//Model .
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {

	dbCfg := config.Cfg.GetStringMapString("database")

	db, err := gorm.Open(dbCfg["type"], fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg["user"],
		dbCfg["password"],
		dbCfg["host"],
		dbCfg["database"]))

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
}

//CloseDB .
func CloseDB() {
	defer db.Close()
}
