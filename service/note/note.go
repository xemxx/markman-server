package note

import (
	"markman-server/model"
)

var db = model.Db

func GetSync(uid, afterSC, count int) ([]model.Note, error) {
	var all []model.Note
	err := db.Where("uid = ? AND SC > ?", uid, afterSC).Order("SC asc").Limit(count).Find(&all).Error
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

func Add(n model.Note) error {
	db.Create(&n)
	return db.Error
}

func Get(guid string) model.Note {
	var note model.Note
	db.Where("guid = ?", guid).First(&note)
	return note
}

func Update(note model.Note) error {
	db.Model(model.Note{}).Where("guid=?", note.Guid).Update(note)
	return db.Error
}
