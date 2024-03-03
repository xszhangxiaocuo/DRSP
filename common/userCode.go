package common

/*用户模块错误代码*/

var (
	CacheError             = NewError(200001, "Cache错误")
	DBError                = NewError(200002, "DB错误")
	NoLegalMobile          = NewError(100001, "手机号不合法")
	CaptchaError           = NewError(100002, "验证码不合法")
	CaptchaNotExist        = NewError(100003, "验证码不存在")
	AccountExist           = NewError(100004, "账号已被注册")
	MobileExist            = NewError(100005, "手机号已被注册")
	AccountOrPasswordError = NewError(100006, "用户或密码错误")
	NoLogin                = NewError(100007, "未登录")
)
