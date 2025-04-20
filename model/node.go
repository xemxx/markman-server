package model

import "time"

// Node 统一的节点结构，可以是文件夹或文件
type Node struct {
	ID         int       `json:"id"`
	Guid       string    `json:"guid"`
	Uid        int       `json:"uid"`
	Type       string    `json:"type"`     // "note" 或 "folder"
	Title      string    `json:"title"`    // 文件夹名称或文件标题
	ParentId   string    `json:"parentId"` // 父节点的Guid，根节点为'root'
	Content    string    `gorm:"type:LONGTEXT" json:"content"`
	Sort       int       `json:"sort"`
	SortType   int       `gorm:"column:sortType" json:"sortType"`
	SC         int       `gorm:"column:SC" json:"SC"`
	AddDate    time.Time `gorm:"column:addDate" json:"addDate"`
	ModifyDate time.Time `gorm:"column:modifyDate" json:"modifyDate"`
	IsDel      int       `gorm:"column:isDel" json:"isDel"`
}

func (u *Node) TableName() string {
	return "node"
}
