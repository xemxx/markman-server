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

func Add(n model.Notebook){
	n.Add()
}

func Get(guid string)model.Notebook{
	var notebook model.Notebook
	db.Where("guid = ?",[]string{guid}).First(&notebook)
	return notebook
}

func Update(notebook model.Notebook){
	db.Model(model.Notebook{}).Where("guid=?",[]string{notebook.Guid}).Update(notebook)
}