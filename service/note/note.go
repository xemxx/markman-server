package note

import (
	"markman-server/model"
)

func GetSync(uid, afterSC, count int) ([]model.Note, error) {
	var all []model.Note
	err := model.I().Where("uid = ? AND SC > ?", uid, afterSC).Order("SC asc").Limit(count).Find(&all).Error
	if err != nil {
		return []model.Note{}, err
	}
	return all, nil
}

func Exist(guid string) int {
	note := model.Note{
		Guid: guid,
	}
	d := model.I().Where(&note).First(&note)
	if note.ID == 0 || d.Error != nil {
		return 0
	}
	return note.ID
}

func Add(n model.Note) error {
	d := model.I().Create(&n)
	return d.Error
}

func Get(guid string) model.Note {
	var note model.Note
	model.I().Where("guid = ?", guid).First(&note)
	return note
}

func Update(note model.Note) error {
	d := model.I().Model(model.Note{}).Where("guid=?", note.Guid).Updates(note)
	return d.Error
}
