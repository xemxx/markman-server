package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"

	"markman-server/service/user"
	"markman-server/tools/common"
	"markman-server/tools/response"
)

// Sign .
type Sign struct {
	Username string `form:"username" json:"username" validate:"required,min=3,max=50"`
	Password string `form:"password" json:"password" validate:"required,min=3,max=50"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// SignUp 注册流程，生成新token并供使用
func SignUp(c *gin.Context) {
	params, err := validateSign(c)
	if err != nil {
		response.JSON(c, response.InvalidParams, response.GetMsg(response.InvalidParams))
		return
	}

	isExist := user.ExistUserByName(params.Username)
	if isExist {
		response.JSON(c, response.ErrorExistUserName, response.GetMsg(response.ErrorExistUserName))
		return
	}
	uid, ok := user.AddUser(params.Username, params.Password)
	if !ok {
		response.JSON(c, response.ErrorInsertFailed, response.GetMsg(response.ErrorInsertFailed))
		return
	}
	response.JSON(c, response.SUCCESS, response.GetMsg(response.SUCCESS), uid)
}

// SignIn 登录流程，生成新token并供使用
func SignIn(c *gin.Context) {
	params, err := validateSign(c.Copy())
	if err != nil {
		response.JSON(c, response.InvalidParams, response.GetMsg(response.InvalidParams))
		return
	}

	var code int
	data := make(map[string]interface{})
	slog.Info(fmt.Sprintln(params))
	info, isExist := user.GetByPass(params.Username, params.Password)
	if !isExist {
		code = response.ErrorUser
		response.JSON(c, code, response.GetMsg(code), data)
		return
	}

	token, err := common.GenerateToken(params.Username, info.ID)
	if err != nil {
		code = response.ERROR
		response.JSON(c, code, response.GetMsg(code), data)
		return
	}

	_ = user.SaveToken(info.ID, token)
	data["token"] = token
	data["uuid"] = info.UUID
	code = response.SUCCESS
	response.JSON(c, code, response.GetMsg(code), data)
}

// 验证登录时和注册提交的参数是否合法
func validateSign(c *gin.Context) (Sign, error) {
	params := Sign{}
	if err := c.ShouldBindJSON(&params); err != nil {
		slog.Error(err.Error())
		return Sign{}, err
	}
	if err := validate.Struct(params); err != nil {
		return Sign{}, err
	}

	return params, nil
}
