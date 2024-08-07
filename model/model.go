package model

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
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
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Database)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	if Db.Migrator().HasColumn(&User{}, "uid") {
		if err = Db.Migrator().DropColumn(&User{}, "uid"); err != nil {
			return err
		}
	}
	if Db.Migrator().HasColumn(&User{}, "create_time") {
		if err = Db.Migrator().RenameColumn(&User{}, "create_time", "created_at"); err != nil {
			return err
		}
	}
	err = Db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	// update default uuid
	tx := Db.Begin()
	users := []User{}
	tx.Select("*").Find(&users)
	for _, user := range users {
		if user.UUID == "" {
			user.UUID = uuid.New().String()
			if user.CreatedAt.IsZero() {
				user.CreatedAt = time.Now()
			}
			tx.Save(&user)
		}
	}
	slog.Debug("users", "len", len(users))
	tx.Commit()

	if err = Db.AutoMigrate(&Note{}); err != nil {
		return err
	}
	return Db.AutoMigrate(&Notebook{})
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
