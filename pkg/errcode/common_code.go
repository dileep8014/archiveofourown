package errcode

var (
	Success                  = NewError(0, "成功")
	ServerError              = NewError(10000000, "服务内部错误")
	InValidParams            = NewError(10000001, "入参错误")
	NotFound                 = NewError(10000002, "资源查找失败,所请求的资源不存在")
	UnauthorizedAuthNotExist = NewError(10000003, "鉴权失败")
	UnauthorizedTokenError   = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeOut = NewError(10000005, "鉴权失败，Token超时")
	TooManyRequest           = NewError(10000007, "请求过多")
	ErrorPermission          = NewError(10000008, "权限不足")
)
