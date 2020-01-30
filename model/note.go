package model

import "time"

type Note struct {
	ID         int       `json:"id"`
	Guid       string    `json:"guid"`
	Uid        int       `json:"uid"`
	Bid        int       `json:"bid"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	SC         int       `json:"SC"`
	AddDate    time.Time `json:"addDate"`
	ModifyDate time.Time `json:"modifyDate"`
	IsDel      int       `json:"isDel"`
}
