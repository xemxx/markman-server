package model

import "time"

type Notebook struct {
	ID         int       `json:"id"`
	Guid       string    `json:"guid"`
	Uid        int       `json:"uid"`
	Name       string    `json:"name"`
	Sort       int       `json:"sort"`
	SortType   int       `json:"sortType"`
	SC         int       `json:"SC"`
	AddDate    time.Time `json:"addDate"`
	ModifyDate time.Time `json:"modifyDate"`
	IsDel      int       `json:"isDel"`
}
