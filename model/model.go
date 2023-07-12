package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"markman-server/tools/config"
)

// Db .
var Db *gorm.DB

func Init() error {
	dbCfg := config.Cfg.GetStringMapString("database")
	var err error
	Db, err = gorm.Open(dbCfg["type"], fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		dbCfg["user"],
		dbCfg["password"],
		dbCfg["host"],
		dbCfg["database"]))
	if err != nil {
		return err
	}
	Db.SingularTable(true)
	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)
	Db.AutoMigrate(&User{}).AutoMigrate(&Note{}).AutoMigrate(&Notebook{})
	return nil
}

// CloseDB .
func CloseDB() {
	defer Db.Close()
}

func I() *gorm.DB {
	return Db
}

// Test .
func Test() {
	user := &User{}
	Db.Select("id")
	Db.Where(&User{Username: "admin"})
	Db.First(user)
	fmt.Println(user)
}
