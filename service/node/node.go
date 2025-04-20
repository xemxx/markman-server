package node

import (
	"markman-server/model"
)

// GetSync 获取同步数据
func GetSync(uid, afterSC, count int) ([]model.Node, error) {
	var all []model.Node
	err := model.I().Where("uid = ? AND SC > ?", uid, afterSC).Order("SC asc").Limit(count).Find(&all).Error
	if err != nil {
		return []model.Node{}, err
	}
	return all, nil
}

// Exist 检查节点是否存在
func Exist(guid string) int {
	node := model.Node{
		Guid: guid,
	}
	d := model.I().Where(&node).First(&node)
	if node.ID == 0 || d.Error != nil {
		return 0
	}
	return node.ID
}

// Add 添加节点
func Add(n model.Node) error {
	d := model.I().Create(&n)
	return d.Error
}

// Get 获取节点
func Get(guid string) model.Node {
	var node model.Node
	model.I().Where("guid = ?", guid).First(&node)
	return node
}

// Update 更新节点
func Update(node model.Node) error {
	d := model.I().Model(model.Node{}).Where("guid=?", node.Guid).Updates(node)
	return d.Error
}
