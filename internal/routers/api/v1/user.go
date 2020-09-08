package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/internal/middleware"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

type User struct {
}

func NewUser() User {
	return User{}
}

func (u User) Router(api gin.IRouter) {
	// 注册
	api.POST("/register", u.Register)
	api.GET("/register/identify", u.Identify)
	api.POST("/register/create", u.CreateUser)
	// 登录
	api.POST("/login", u.Login)
	// 退出登录
	api.POST("/logout", middleware.Verify(middleware.LoginNeed), u.Logout)
	// 获取当前登录用户信息
	api.GET("/currentUser", middleware.Verify(middleware.LoginNeed), u.CurrentUser)
	// 当前登录人偏好设置
	api.GET("/currentUser/setting", middleware.Verify(middleware.LoginNeed), u.CurrentSetting)
	// 修改用户信息
	api.POST("/currentUser", middleware.Verify(middleware.LoginNeed), u.Update)
	// 修改用户邮箱
	api.POST("/currentUser/email", middleware.Verify(middleware.LoginNeed), u.UpdateEmail)
	// 修改用户密码
	api.POST("/currentUser/password", middleware.Verify(middleware.LoginNeed), u.UpdatePass)
	// 修改用户偏好设置信息
	api.POST("/currentUser/setting", middleware.Verify(middleware.LoginNeed), u.UpdateSetting)
	// 获取指定用户信息
	api.GET("/user/:id", u.UserInfo)
}

// @Summary 注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param email query string true "注册邮箱" email
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/register [post]
func (u User) Register(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.UserRegisterRequest{}
	valid, errors := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errors.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err := svc.RegisterUser(req.Email)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorRegisterUser)
		return
	}
	res.ToSuccessResponse()
}

// @Summary 验证注册邮箱
// @Tags 用户
// @Accept json
// @Produce json
// @Param path query string true "注册路径"
// @Success 200 {string} string "邮箱"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/register/identify [get]
func (u User) Identify(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	path := ctx.Query("path")
	if path == "" {
		res.ToErrorResponse(errcode.InValidParams.WithDetails("参数path为空"))
		return
	}
	svc := service.NewService(ctx)
	email, err := svc.Identify(path)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorIdentifyUser)
		return
	}
	res.ToResponse(email)
}

// @Summary 创建用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body service.UserCreateRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/register/create [post]
func (u User) CreateUser(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.UserCreateRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err := svc.CreateUser(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorCreateUser)
		return
	}
	res.ToSuccessResponse()
}

// @Summary 登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body service.UserLoginRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/login [post]
func (u User) Login(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.UserLoginRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err := svc.Login(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUserLogin)
		return
	}
	res.ToSuccessResponse()
}

// @Summary 退出登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/logout [post]
func (u User) Logout(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)
	svc.Logout()
	res.ToSuccessResponse()
}

// @Summary 当前用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} service.UserResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/currentUser [get]
func (u User) CurrentUser(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)
	user, err := svc.UserInfo(ctx.GetInt64("me.id"))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUserInfo)
		return
	}
	res.ToResponse(user)
}

// @Summary 当前用户偏好设置
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} service.UserSettingResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/currentUser/setting [get]
func (u User) CurrentSetting(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)
	st, err := svc.CurrentUserSetting()
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUserSetting)
		return
	}
	res.ToResponse(st)
}

// @Summary 修改用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body service.UserUpdateRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/currentUser [post]
func (u User) Update(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.UserUpdateRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err := svc.UpdateUserInfo(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUserUpdate)
		return
	}
	res.ToSuccessResponse()
}

// @Summary 修改用户邮箱
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body service.UserUpdateEmailRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/currentUser/email [post]
func (u User) UpdateEmail(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.UserUpdateEmailRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err := svc.UpdateUserEmail(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUserUpdate)
		return
	}
	res.ToSuccessResponse()
}

// @Summary 修改用户密码
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body service.UserUpdatePassRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/currentUser/password [post]
func (u User) UpdatePass(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.UserUpdatePassRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err := svc.UpdateUserPassword(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUserUpdate)
		return
	}
	res.ToSuccessResponse()
}

// @Summary 修改用户偏好设置
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body service.UserSettingRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/currentUser/setting [post]
func (u User) UpdateSetting(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.UserSettingRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err := svc.UpdateUserSetting(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUserUpdateSetting)
		return
	}
	res.ToSuccessResponse()
}

// @Summary 用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} service.UserResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user/{id} [get]
func (u User) UserInfo(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	id, err := app.ShouldParamConvertInt(ctx, "id")
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}
	svc := service.NewService(ctx)
	user, err := svc.UserInfo(int64(id))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUserInfo)
		return
	}
	res.ToResponse(user)
}
