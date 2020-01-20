package user

import (
	"github.com/go-playground/validator/v10"
	"markman-server/service/user"
	"markman-server/tools/common"
	"markman-server/tools/logs"
	"markman-server/tools/response"

	"github.com/gin-gonic/gin"
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
	if !user.AddUser(params.Username, params.Password) {
		response.JSON(c, response.ErrorInsertFailed, response.GetMsg(response.ErrorInsertFailed))
		return
	}
	response.JSON(c, response.SUCCESS, response.GetMsg(response.SUCCESS))
}

// SignIn 登录流程，生成新token并供使用
func SignIn(c *gin.Context) {
	params, err := validateSign(c.Copy())
	if err != nil {
		response.JSON(c, response.InvalidParams, response.GetMsg(response.InvalidParams))
		return
	}

	code := response.SUCCESS
	data := make(map[string]interface{})
	uid, isExist := user.ExistUser(params.Username, params.Password)
	if !isExist {
		code = response.ErrorUser
		response.JSON(c, code, response.GetMsg(code), data)
		return
	}

	token, err := common.GenerateToken(params.Username, uid)
	if err != nil {
		code = response.ErrorAuthToken
		response.JSON(c, code, response.GetMsg(code), data)
		return
	}

	_ = user.SaveToken(uid, token)
	data["token"] = token
	code = response.SUCCESS
	response.JSON(c, code, response.GetMsg(code), data)
}

// 验证登录时和注册提交的参数是否合法
func validateSign(c *gin.Context) (Sign, error) {
	params := Sign{}
	if err := c.ShouldBind(&params); err != nil {
		logs.Log(err.Error())
		return Sign{}, err
	}
	if err := validate.Struct(params); err != nil {
		return Sign{}, err
	}

	return params, nil
}
