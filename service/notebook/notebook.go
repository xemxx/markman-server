package notebook

import (
	"markman-server/model"
)

var db = model.Db

func GetSync(uid, afterSC, count int) ([]model.Notebook, error) {
	var all []model.Notebook
	err := db.Where("uid = ? AND SC > ?", uid, afterSC).Order("SC asc").Limit(count).Find(&all).Error
	if err != nil {
		return []model.Notebook{}, err
	}
	return all, nil
}

func Exist(guid string) int {
	notebook := model.Notebook{
		Guid: guid,
	}
	db.Where(&notebook).First(&notebook)
	if notebook.ID == 0 || db.Error != nil {
		return 0
	}
	return notebook.ID
}

func Add(n model.Notebook) error {
	db.Create(&n)
	return db.Error
}

func Get(guid string) model.Notebook {
	var notebook model.Notebook
	db.Where("guid = ?", guid).First(&notebook)
	return notebook
}

func Update(notebook model.Notebook) error {
	db.Model(model.Notebook{}).Where("guid=?", notebook.Guid).Update(notebook)
	return db.Error
}
