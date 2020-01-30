package notebook

import (
	"markman-server/model"
)

var db = model.Db

func GetSync(uid, afterSC, count int) ([]model.Notebook, error) {
	var all []model.Notebook
	err := db.Where("uid = ? AND SC > ?", []int{uid, afterSC}).Order("SC aes").Limit(count).Scan(&all).Error
	if err != nil {
		return []model.Notebook{}, err
	}
	return all, nil
}
