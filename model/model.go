package model

import (
	"fmt"
	"log/slog"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"markman-server/tools/config"
)

// Db .
var Db *gorm.DB

func Init() error {
	dbCfg := config.Cfg.Database
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Database)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		slog.Error("connect database failed", "err", err, "dsn", dsn)
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
			id, err := gonanoid.New()
			if err != nil {
				slog.Error("BUG: generate uuid failed", "err", err)
				return err
			}
			user.UUID = id
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
	if err = Db.AutoMigrate(&Notebook{}); err != nil {
		return err
	}
	if err = Db.AutoMigrate(&Node{}); err != nil {
		return err
	}

	// 迁移数据
	return migrateDataToNode()
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

// migrateDataToNode 将note和notebook数据迁移到node表中
func migrateDataToNode() error {
	// 检查是否已经迁移过数据
	var count int64
	Db.Model(&Node{}).Count(&count)
	if count > 0 {
		slog.Info("Node表已有数据，跳过迁移")
		return nil
	}

	// 开始事务
	tx := Db.Begin()

	// 迁移notebook数据
	var notebooks []Notebook
	if err := tx.Find(&notebooks).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("获取notebook数据失败: %w", err)
	}

	slog.Info("开始迁移notebook数据", "count", len(notebooks))
	for _, nb := range notebooks {
		node := Node{
			Guid:       nb.Guid,
			Uid:        nb.Uid,
			Type:       "folder",
			Title:      nb.Name,
			ParentId:   "root", // notebook默认为根节点
			Content:    "",
			Sort:       nb.Sort,
			SortType:   nb.SortType,
			SC:         nb.SC,
			AddDate:    nb.AddDate,
			ModifyDate: nb.ModifyDate,
			IsDel:      nb.IsDel,
		}

		if err := tx.Create(&node).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("创建notebook对应的node失败: %w", err)
		}
	}

	// 迁移note数据
	var notes []Note
	if err := tx.Find(&notes).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("获取note数据失败: %w", err)
	}

	slog.Info("开始迁移note数据", "count", len(notes))
	for _, n := range notes {
		node := Node{
			Guid:       n.Guid,
			Uid:        n.Uid,
			Type:       "note",
			Title:      n.Title,
			ParentId:   n.Bid, // note的Bid字段对应node的ParentId
			Content:    n.Content,
			Sort:       0, // note没有排序字段，默认为0
			SortType:   0, // note没有排序类型字段，默认为0
			SC:         n.SC,
			AddDate:    n.AddDate,
			ModifyDate: n.ModifyDate,
			IsDel:      n.IsDel,
		}

		if err := tx.Create(&node).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("创建note对应的node失败: %w", err)
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	slog.Info("数据迁移完成", "notebooks", len(notebooks), "notes", len(notes))
	return nil
}
