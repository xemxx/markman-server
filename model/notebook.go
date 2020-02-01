package model

import "time"

type Notebook struct {
	ID         int       `json:"id"`
	Guid       string    `json:"guid"`
	Uid        int       `json:"uid"`
	Name       string    `json:"name"`
	Sort       int       `json:"sort"`
	SortType   int       `gorm:"column:sortType" json:"sortType"`
	SC         int       `gorm:"column:SC" json:"SC"`
	AddDate    time.Time `gorm:"column:addDate" json:"addDate"`
	ModifyDate time.Time `gorm:"column:modifyDate" json:"modifyDate"`
	IsDel      int       `gorm:"column:isDel" json:"isDel"`
}
