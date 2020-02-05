package notebook

import (
	"github.com/gin-gonic/gin"
	"log"
	"markman-server/model"
	"markman-server/service/notebook"
	"markman-server/service/user"
	"markman-server/tools/response"
	"strconv"
	"time"
)

type Client struct {
	ID          int    `json:"id"`
	Guid        string `json:"guid"`
	Uid         int    `json:"uid"`
	Name        string `json:"name"`
	Sort        int    `json:"sort"`
	SortType    int    `json:"sortType"`
	ModifyState int    `json:"modifyState"`
	SC          int    `json:"SC"`
	AddDate     int64  `json:"addDate"`
	ModifyDate  int64  `json:"modifyDate"`
}

func GetSync(c *gin.Context) {
	afterSC, _ := strconv.Atoi(c.DefaultQuery("afterSC", "0"))
	maxCount, _ := strconv.Atoi(c.DefaultQuery("maxCount", "10"))
	uid := c.GetInt("uid")

	code, data := response.SUCCESS, make(map[string]interface{})
	notebooks, err := notebook.GetSync(uid, afterSC, maxCount)
	if err != nil {
		log.Println(err)
		code = response.ERROR
	} else {
		data["notebooks"] = notebooks
	}
	response.JSON(c, code, response.GetMsg(code), data)
}

type resultErr struct {
	IsRepeat bool `json:"isRepeat"`
	IsErr    bool `json:"isErr"`
	SC       int  `json:"SC"`
}

func Create(c *gin.Context) {
	var client Client
	_ = c.ShouldBind(&client)
	uid := c.GetInt("uid")
	u := user.Get(uid)

	code, data := response.SUCCESS, resultErr{}

	id := notebook.Exist(client.Guid)
	if id == 0 {
		newNotebook := model.Notebook{
			Guid:       client.Guid,
			Uid:        c.GetInt("uid"),
			Name:       client.Name,
			Sort:       client.Sort,
			SortType:   client.SortType,
			SC:         u.SC + 1,
			AddDate:    time.Unix(client.AddDate, 0),
			ModifyDate: time.Unix(client.ModifyDate, 0),
			IsDel:      0,
		}
		err := notebook.Add(newNotebook)
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

func Delete(c *gin.Context) {
	var client Client
	_ = c.ShouldBind(&client)
	uid := c.GetInt("uid")
	u := user.Get(uid)

	code, data := response.SUCCESS, resultErr{}

	local := notebook.Get(client.Guid)
	if local.SC == client.SC {
		newNotebook := model.Notebook{
			Uid:        c.GetInt("uid"),
			Guid:       client.Guid,
			Name:       client.Name,
			Sort:       client.Sort,
			SortType:   client.SortType,
			SC:         u.SC + 1,
			AddDate:    time.Unix(client.AddDate, 0),
			ModifyDate: time.Unix(client.ModifyDate, 0),
			IsDel:      1,
		}
		err := notebook.Update(newNotebook)
		if err != nil {
			data = resultErr{false, true, u.SC + 1}
		} else {
			user.UpdateSC(uid, u.SC+1)
			data = resultErr{false, false, u.SC + 1}
		}
	} else {
		data = resultErr{
			IsRepeat: false,
			IsErr:    true,
			SC:       u.SC,
		}
	}
	response.JSON(c, code, response.GetMsg(code), data)
}
func Update(c *gin.Context) {
	var client Client
	_ = c.ShouldBind(&client)
	uid := c.GetInt("uid")
	u := user.Get(uid)

	code, data := response.SUCCESS, resultErr{}

	local := notebook.Get(client.Guid)
	if local.SC == client.SC {
		newNotebook := model.Notebook{
			Guid:       client.Guid,
			Uid:        c.GetInt("uid"),
			Name:       client.Name,
			Sort:       client.Sort,
			SortType:   client.SortType,
			SC:         u.SC + 1,
			AddDate:    time.Unix(client.AddDate, 0),
			ModifyDate: time.Unix(client.ModifyDate, 0),
		}
		err := notebook.Update(newNotebook)
		if err != nil {
			data = resultErr{false, true, u.SC + 1}
		} else {
			user.UpdateSC(uid, u.SC+1)
			data = resultErr{false, false, u.SC + 1}
		}
	} else {
		data = resultErr{
			IsRepeat: false,
			IsErr:    true,
			SC:       u.SC,
		}
	}
	response.JSON(c, code, response.GetMsg(code), data)
}
