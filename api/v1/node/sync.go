package node

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"markman-server/model"
	"markman-server/service/node"
	"markman-server/service/user"
	"markman-server/tools/response"
)

// Client 客户端传来的节点数据
type Client struct {
	ID          int    `json:"id"`
	Guid        string `json:"guid"`
	Uid         int    `json:"uid"`
	Type        string `json:"type"`        // "note" 或 "folder"
	Title       string `json:"title"`       // 文件夹名称或文件标题
	ParentId    string `json:"parentId"`    // 父节点的guid，根节点为'root'
	Content     string `json:"content"`     // 文件内容
	Sort        int    `json:"sort"`        // 排序
	SortType    int    `json:"sortType"`    // 排序类型
	ModifyState int    `json:"modifyState"` // 修改状态
	SC          int    `json:"SC"`          // 同步计数
	AddDate     int64  `json:"addDate"`     // 添加日期
	ModifyDate  int64  `json:"modifyDate"`  // 修改日期
}

// NodeResponse 返回给前端的节点数据
type NodeResponse struct {
	ID         int    `json:"id"`
	Guid       string `json:"guid"`
	Uid        int    `json:"uid"`
	Type       string `json:"type"`       // "note" 或 "folder"
	Title      string `json:"title"`      // 文件夹名称或文件标题
	ParentId   string `json:"parentId"`   // 父节点的guid，根节点为'root'
	Content    string `json:"content"`    // 文件内容
	Sort       int    `json:"sort"`       // 排序
	SortType   int    `json:"sortType"`   // 排序类型
	SC         int    `json:"SC"`         // 同步计数
	AddDate    int64  `json:"addDate"`    // 添加日期
	ModifyDate int64  `json:"modifyDate"` // 修改日期
	IsDel      int    `json:"isDel"`      // 是否删除
}

// @Summary	GetSync Node
// @Schemes
// @Description	Node Sync
// @Tags			node
// @Accept			json
// @Produce		json
// @Param			Authorization	header		string									true	"JWT"
// @Param			afterSC			query		string									0	"同步版本号"
// @Param			maxCount		query		integer									10	"最大个数"
// @Success		200				{object}	response.Response{data=[]model.Node}	"desc"
// @Failure		400				{object}	response.Response						"desc"
// @Router			/node/getSync [get]
func GetSync(c *gin.Context) {
	afterSC, _ := strconv.Atoi(c.DefaultQuery("afterSC", "0"))
	maxCount, _ := strconv.Atoi(c.DefaultQuery("maxCount", "10"))
	uid := c.GetInt("uid")

	code, data := response.SUCCESS, make(map[string]interface{})
	nodes, err := node.GetSync(uid, afterSC, maxCount)
	if err != nil {
		log.Println(err)
		code = response.ERROR
	} else {
		// 将数据库模型转换为响应模型
		responseNodes := make([]NodeResponse, len(nodes))
		for i, n := range nodes {
			responseNodes[i] = NodeResponse{
				ID:         n.ID,
				Guid:       n.Guid,
				Uid:        n.Uid,
				Type:       n.Type,
				Title:      n.Title,
				ParentId:   n.ParentId,
				Content:    n.Content,
				Sort:       n.Sort,
				SortType:   n.SortType,
				SC:         n.SC,
				AddDate:    n.AddDate.Unix(),
				ModifyDate: n.ModifyDate.Unix(),
				IsDel:      n.IsDel,
			}
		}
		data["nodes"] = responseNodes
	}
	response.JSON(c, code, response.GetMsg(code), data)
}

type resultErr struct {
	IsRepeat bool `json:"isRepeat"`
	IsErr    bool `json:"isErr"`
	SC       int  `json:"SC"`
}

// Create 创建节点
func Create(c *gin.Context) {
	var client Client
	_ = c.ShouldBind(&client)
	uid := c.GetInt("uid")
	u := user.Get(uid)

	var data resultErr
	code := response.SUCCESS

	id := node.Exist(client.Guid)
	if id == 0 {
		newNode := model.Node{
			Guid:       client.Guid,
			Uid:        uid,
			Type:       client.Type,
			Title:      client.Title,
			ParentId:   client.ParentId,
			Content:    client.Content,
			Sort:       client.Sort,
			SortType:   client.SortType,
			SC:         u.SC + 1,
			AddDate:    time.Unix(client.AddDate, 0),
			ModifyDate: time.Unix(client.ModifyDate, 0),
			IsDel:      0,
		}
		err := node.Add(newNode)
		if err != nil {
			data = resultErr{false, true, u.SC + 1}
		} else {
			user.UpdateSC(uid, u.SC+1)
			data = resultErr{false, false, u.SC + 1}
		}
	} else {

		data = resultErr{true, false, u.SC}
	}
	response.JSON(c, code, response.GetMsg(code), data)
}

// Delete 删除节点
func Delete(c *gin.Context) {
	var client Client
	_ = c.ShouldBind(&client)
	uid := c.GetInt("uid")
	u := user.Get(uid)

	var data resultErr
	code := response.SUCCESS

	local := node.Get(client.Guid)
	if local.SC == client.SC {
		newNode := model.Node{
			Guid:       client.Guid,
			Uid:        uid,
			Type:       client.Type,
			Title:      client.Title,
			ParentId:   client.ParentId,
			Content:    client.Content,
			Sort:       client.Sort,
			SortType:   client.SortType,
			SC:         u.SC + 1,
			AddDate:    time.Unix(client.AddDate, 0),
			ModifyDate: time.Unix(client.ModifyDate, 0),
			IsDel:      1,
		}
		err := node.Update(newNode)
		if err != nil {
			data = resultErr{false, true, u.SC + 1}
		} else {
			user.UpdateSC(uid, u.SC+1)
			data = resultErr{false, false, u.SC + 1}
		}
	} else {
		log.Printf("error: SC not match, local SC: %d, client SC: %d\n", local.SC, client.SC)
		data = resultErr{
			IsRepeat: false,
			IsErr:    true,
			SC:       u.SC,
		}
	}
	response.JSON(c, code, response.GetMsg(code), data)
}

// Update 更新节点
func Update(c *gin.Context) {
	var client Client
	_ = c.ShouldBind(&client)
	uid := c.GetInt("uid")
	u := user.Get(uid)

	var data resultErr
	code := response.SUCCESS

	local := node.Get(client.Guid)
	if local.SC == client.SC {
		newNode := model.Node{
			Guid:       client.Guid,
			Uid:        uid,
			Type:       client.Type,
			Title:      client.Title,
			ParentId:   client.ParentId,
			Content:    client.Content,
			Sort:       client.Sort,
			SortType:   client.SortType,
			SC:         u.SC + 1,
			AddDate:    time.Unix(client.AddDate, 0),
			ModifyDate: time.Unix(client.ModifyDate, 0),
		}
		err := node.Update(newNode)
		if err != nil {
			log.Println(err)
			data = resultErr{false, true, u.SC + 1}
		} else {
			user.UpdateSC(uid, u.SC+1)
			data = resultErr{false, false, u.SC + 1}
		}
	} else {
		log.Printf("error: SC not match, local SC: %d, client SC: %d\n", local.SC, client.SC)
		data = resultErr{
			IsRepeat: false,
			IsErr:    true,
			SC:       u.SC,
		}
	}
	response.JSON(c, code, response.GetMsg(code), data)
}
