package notebook

import (
	"markman-server/model"
)

func GetSync(uid, afterSC, count int) ([]model.Notebook, error) {
	var all []model.Notebook
	err := model.I().Where("uid = ? AND SC > ?", uid, afterSC).Order("SC asc").Limit(count).Find(&all).Error
	if err != nil {
		return []model.Notebook{}, err
	}
	return all, nil
}

func Exist(guid string) int {
	notebook := model.Notebook{
		Guid: guid,
	}
	d := model.I().Where(&notebook).First(&notebook)
	if notebook.ID == 0 || d.Error != nil {
		return 0
	}
	return notebook.ID
}

func Add(n model.Notebook) error {
	d := model.I().Create(&n)
	return d.Error
}

func Get(guid string) model.Notebook {
	var notebook model.Notebook
	model.I().Where("guid = ?", guid).First(&notebook)
	return notebook
}

func Update(notebook model.Notebook) error {
	d := model.I().Model(model.Notebook{}).Where("guid=?", notebook.Guid).Update(notebook)
	return d.Error
}
