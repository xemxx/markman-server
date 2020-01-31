package note

import (
	"markman-server/model"
)

var db = model.Db

func GetSync(uid, afterSC, count int) ([]model.Note, error) {
	var all []model.Note
	err := db.Where("uid = ? AND SC > ?", []int{uid, afterSC}).Order("SC aes").Limit(count).Find(&all).Error
	if err != nil {
		return []model.Note{}, err
	}
	return all, nil
}

func Exist(guid string) int {
	note := model.Note{
		Guid: guid,
	}
	db.Where(&note).First(&note)
	if note.ID == 0 || db.Error != nil {
		return 0
	}
	return note.ID
}

func Add(n model.Note) {
	db.Create(n)
}

func Get(guid string) model.Note {
	var note model.Note
	db.Where("guid = ?", []string{guid}).First(&note)
	return note
}

func Update(note model.Note) {
	db.Model(model.Note{}).Where("guid=?", []string{note.Guid}).Update(note)
}
