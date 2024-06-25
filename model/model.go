package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"markman-server/tools/config"
)

// Db .
var Db *gorm.DB

func Init() error {
	dbCfg := config.Cfg.Database
	var err error
	Db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		dbCfg.User,
		dbCfg.Type,
		dbCfg.Host,
		dbCfg.Database)), &gorm.Config{})
	if err != nil {
		return err
	}
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Note{})
	Db.AutoMigrate(&Notebook{})
	return nil
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
