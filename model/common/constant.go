package common

const (
	EmailReg = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 正则：邮箱
	PhoneReg = `1\d{10}`                                     // 正则：手机号

	YesStatus   = 1 // 状态：启用、成功、菜单
	NoStatus    = 2 // 状态：禁用、失败、按钮
	OtherStatus = 3 // 状态：未知、过程中

	AccessTokenCache  = "AccessToken"  // 缓存：AccessToken
	RefreshTokenCache = "RefreshToken" // 缓存：RefreshToken
	ManagerTokenCache = "ManagerToken" // 缓存：管理页面token
	ManagerVcodeCache = "ManagerVcode" // 缓存：管理页面验证码
)
