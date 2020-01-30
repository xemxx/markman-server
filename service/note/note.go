package note

import (
	"markman-server/model"
)

var db = model.Db

func GetSync(uid, afterSC, count int) ([]model.Note, error) {
	var all []model.Note
	err := db.Where("uid = ? AND SC > ?", []int{uid, afterSC}).Order("SC aes").Limit(count).Scan(&all).Error
	if err != nil {
		return []model.Note{}, err
	}
	return all, nil
}
