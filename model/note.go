package model

import "time"

type Note struct {
	ID         int       `json:"id"`
	Guid       string    `json:"guid"`
	Uid        int       `json:"uid"`
	Bid        string    `json:"bid"`
	Title      string    `json:"title"`
	Content    string    `gorm:"type:LONGTEXT" json:"content"`
	SC         int       `gorm:"column:SC" json:"SC"`
	AddDate    time.Time `gorm:"column:addDate" json:"addDate"`
	ModifyDate time.Time `gorm:"column:modifyDate" json:"modifyDate"`
	IsDel      int       `gorm:"column:isDel" json:"isDel"`
}
