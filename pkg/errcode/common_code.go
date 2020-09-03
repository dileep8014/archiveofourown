package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InValidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "资源查找失败")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeOut  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequest            = NewError(10000007, "请求过多")
	NotExist                  = NewError(10000008, "资源不存在")
	ErrorPermission           = NewError(10000009, "权限不足")
	VerifyError               = NewError(10000010, "权限校验错误")
)
