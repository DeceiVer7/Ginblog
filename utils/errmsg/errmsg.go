package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code= 2000... 文章模块的错误

	ERROR_ART_NOT_EXIST = 2001
	// code= 3000... 分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002

	// code= 4000... 点赞模块的错误
	ERROR_LIKE             = 4001
	ERROR_CANCELLIKE       = 4002
	ERROR_LIKE_COUNT       = 4003
	ERROR_CANCELLIKE_COUNT = 4004
	ERROR_LIKE_LIST        = 4005

	// code= 5000... 点赞模块的错误
	ERROR_STAR             = 5001
	ERROR_CANCELSTAR       = 5002
	ERROR_STAR_COUNT       = 5003
	ERROR_CANCELSTAR_COUNT = 5004
	ERROR_STAR_LIST        = 5005
)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_ART_NOT_EXIST: "文章不存在",

	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",

	ERROR_LIKE:             "点赞失败",
	ERROR_CANCELLIKE:       "取消点赞失败",
	ERROR_LIKE_COUNT:       "点赞数量增加失败",
	ERROR_CANCELLIKE_COUNT: "点赞数量减少失败",
	ERROR_LIKE_LIST:        "点赞列表查询失败",

	ERROR_STAR:             "收藏失败",
	ERROR_CANCELSTAR:       "取消收藏失败",
	ERROR_STAR_COUNT:       "收藏数量增加失败",
	ERROR_CANCELSTAR_COUNT: "收藏数量减少失败",
	ERROR_STAR_LIST:        "收藏列表查询失败",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
