package response

const (
	//SUCCESS .
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	//ERROR_EXIST_TAG         = 10001
	//ERROR_NOT_EXIST_TAG     = 10002
	//ERROR_NOT_EXIST_ARTICLE = 10003

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004
	ErrorUser                  = 20005
	ErrorExistUserName         = 20006
	ErrorNotExistUserName      = 20007

	ErrorInsertFailed = 30001
)

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	//ERROR_EXIST_TAG:                "已存在该标签名称",
	//ERROR_NOT_EXIST_TAG:            "该标签不存在",
	//ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",

	ErrorAuthCheckTokenFail:    "Token无效",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token参数错误",
	ErrorUser:                  "账号或者密码错误",
	ErrorExistUserName:         "用户名已存在",
	ErrorNotExistUserName:      "用户名不存在",

	ErrorInsertFailed: "数据库插入失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
