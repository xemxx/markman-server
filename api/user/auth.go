package user

import (
	"log"
	"markman-server/service/user"
	"markman-server/tools/e"
	"markman-server/tools/jwt"
	"markman-server/tools/response"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// Sign .
type Sign struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=3,max=50"`
}

var validate *validator.Validate

// SignUp .
func SignUp(c *gin.Context) {

	params, err := validateSign(c)
	if err != nil {
		response.JSON(c, e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS))
		return
	}

	isExist := user.ExistUserByName(params.Username)
	if isExist {
		response.JSON(c, e.ERROR, "用户名已存在")
		return
	}
	user.AddUser(params.Username, params.Password)
	response.JSON(c, e.SUCCESS, e.GetMsg(e.SUCCESS))
}

// SignIn .
func SignIn(c *gin.Context) {
	params, err := validateSign(c.Copy())
	if err != nil {
		response.JSON(c, e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS))
		return
	}

	code := e.SUCCESS
	data := make(map[string]interface{})
	uid, isExist := user.ExistUser(params.Username, params.Password)
	if isExist {
		token, err := jwt.GenerateToken(params.Username, uid)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			user.SaveToken(uid)
			data["token"] = token
			code = e.SUCCESS
		}
	} else {
		code = e.ERROR_USER
	}
	response.JSON(c, code, e.GetMsg(code), data)
}

func validateSign(c *gin.Context) (Sign, error) {
	var params Sign
	//TODO:fix error: invalid character 'u' looking for beginning of value
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Println(err)
		return Sign{}, err
	}
	if err := validate.Struct(params); err != nil {
		// defer func(err error) {
		// 	for _, err := range err.(validator.ValidationErrors) {
		// 		fmt.Printf("Validate faild: Value %v  and Field %v \n", err.Value(), err.Field())
		// 	}
		// }(err)
		return Sign{}, err
	}

	return params, nil
}
