package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/internal/cache"
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/internal/mq"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"golang.org/x/crypto/bcrypt"
	"net/url"
	"time"
)

// User register request
type UserRegisterRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// User create request
type UserCreateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

// User login request
type UserLoginRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe"`
}

// User update request
type UserUpdateRequest struct {
	Username  string `json:"username" binding:"max=20"`
	Gender    *int   `json:"gender"`
	Introduce string `json:"introduce"`
	Avatar    string `json:"avatar"`
}

// User update email request
type UserUpdateEmailRequest struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// User update password request
type UserUpdatePassRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	Password    string `json:"password" binding:"required,min=8,max=20"`
}

// User update setting request
type UserSettingRequest struct {
	ShowEmail         *bool `json:"showEmail" binding:"required"`
	DisableSearch     *bool `json:"disableSearch" binding:"required"`
	ShowAdult         *bool `json:"showAdult" binding:"required"`
	HiddenGrade       *bool `json:"hiddenGrade" binding:"required"`
	HiddenTag         *bool `json:"hiddenTag" binding:"required"`
	SubscriptionEmail *bool `json:"subscriptionEmail" binding:"required"`
	TopicEmail        *bool `json:"topicEmail" binding:"required"`
	CommentEmail      *bool `json:"commentEmail" binding:"required"`
	SystemEmail       *bool `json:"systemEmail" binding:"required"`
}

// User response example
type UserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Root      bool      `json:"root"`
	Avatar    string    `json:"avatar"`
	Gender    int       `json:"gender"`
	Introduce string    `json:"introduce"`
	WorksNums int64     `json:"worksNums"`
	WorkDay   int64     `json:"workDay"`
	Words     int64     `json:"words"`
	FansNums  int64     `json:"fansNums"`
	CreatedAt time.Time `json:"createdAt"`
}

// User setting response
type UserSettingResponse struct {
	ShowEmail         bool `json:"showEmail" binding:"required"`
	DisableSearch     bool `json:"disableSearch" binding:"required"`
	ShowAdult         bool `json:"showAdult" binding:"required"`
	HiddenGrade       bool `json:"hiddenGrade" binding:"required"`
	HiddenTag         bool `json:"hiddenTag" binding:"required"`
	SubscriptionEmail bool `json:"subscriptionEmail" binding:"required"`
	TopicEmail        bool `json:"topicEmail" binding:"required"`
	CommentEmail      bool `json:"commentEmail" binding:"required"`
	SystemEmail       bool `json:"systemEmail" binding:"required"`
}

func (u UserSettingResponse) TableName() string {
	return model.UserSt{}.TableName()
}

// RegisterUser: 用户注册
func (svc Service) RegisterUser(email string) (err error) {
	defer errwrap.Add(&err, "service.RegisterUser")

	// 检查邮箱是否已注册
	var count int64
	err = svc.db.Model(model.User{}).Select("*").Where("email=?", email).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	// 若已存在，返回错误
	if count > 0 {
		return errcode.ErrorEmailExist
	}
	// 查询是否已经有待验证邮件
	var identify = model.Identify{Email: email}
	err = svc.db.First(&identify, "email=?", identify.Email).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		// 生成注册路径
		var path []byte
		path, err = bcrypt.GenerateFromPassword([]byte(email), 10)
		if err != nil {
			return
		}
		// 存入验证表
		identify.Path = string(path)
		err = svc.db.Create(&identify).Error
		if err != nil {
			return
		}
	}
	// 发送邮件
	go mq.RegisterProvider{}.Send(
		mq.RegisterMsg{
			Email: email,
			Path:  "http://localhost:8000/register/" + url.QueryEscape(identify.Path),
		},
		time.NewTimer(time.Minute))
	return
}

// Identify: 用户注册验证
func (svc Service) Identify(path string) (email string, err error) {
	defer errwrap.Add(&err, "service.Identify")

	// 查询是否有待验证邮件
	path, err = url.PathUnescape(path)
	if err != nil {
		return
	}
	var identify model.Identify
	result := svc.db.Select("email").First(&identify, "path=?", path)
	email, err = identify.Email, CheckError(result, Select_OP)
	return
}

// CreateUser: 创建用户
func (svc Service) CreateUser(req UserCreateRequest) (err error) {
	defer errwrap.Add(&err, "service.CreateUser")

	err = svc.db.Transaction(func(tx *gorm.DB) error {
		// 查询邮箱是否已存在
		var count int64
		err := tx.Model(model.User{}).Select("*").Where("email = ?", req.Email).Count(&count).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if count > 0 {
			return errcode.ErrorEmailExist
		}
		// 查询用户名是否已存在
		err = tx.Model(model.User{}).Select("*").Where("username = ?", req.Username).Count(&count).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if count > 0 {
			return errcode.ErrorUsernameExist
		}
		// 创建账户
		pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			return err
		}
		user := model.User{
			Username: req.Username,
			Email:    req.Email,
			Password: string(pass),
		}
		result := tx.Create(&user)
		err = CheckError(result, Insert_OP)
		if err != nil {
			return err
		}
		result = tx.Create(&model.UserEx{
			UserId: user.ID,
		})
		err = CheckError(result, Insert_OP)
		if err != nil {
			return err
		}
		result = tx.Create(&model.UserSt{
			UserId:      user.ID,
			SystemEmail: true,
		})
		err = CheckError(result, Insert_OP)
		if err != nil {
			return err
		}
		// 存入缓存
		userCache := &cache.UserCache{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
		}
		err = userCache.HMSetAll()
		if err != nil {
			return err
		}
		// 删除待验证表
		result = tx.Delete(&model.Identify{}, "email=?", req.Email)
		return CheckError(result, Delete_OP)
	})
	return
}

// Login: 登录
func (svc Service) Login(req UserLoginRequest) (err error) {
	defer errwrap.Wrap(&err, "service.Login")

	// 验证用户是否存在
	user := model.User{}
	err = svc.db.Select("id,username,root,password").First(&user, "email = ? or username = ?", req.Username, req.Username).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return errcode.ErrorUserNotExist
	}
	if err != nil {
		return
	}
	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return errcode.ErrorUserPassword
	}
	if err != nil {
		return
	}
	// 生成token
	expire := time.Duration(global.JWTSetting.Expire) * time.Second
	if req.RememberMe {
		expire = time.Hour * 24 * 7
	}
	token, err := app.GenerateToken(user.ID, user.Username, user.Root, expire)
	if err != nil {
		return
	}
	// 存入redis
	err = cache.Token{V: token}.SetEX(int(expire.Seconds()))
	if err != nil {
		return
	}
	svc.ctx.Header("Authorization", token)
	return
}

// Logout: 退出登录
func (svc Service) Logout() {
	token := &cache.Token{V: svc.ctx.GetHeader("Authorization")}
	token.Del()
}

// UserInfo: 获取用户信息
func (svc Service) UserInfo(id int64) (resp UserResponse, err error) {
	defer errwrap.Wrap(&err, "service.UserInfo")

	// 从redis获取
	userCache := &cache.UserCache{ID: id}
	err = userCache.HGetAll()
	if err != nil && err == redis.ErrNil {
		// 记录日志
		logger := svc.ctx.Value("logger").(zerolog.Logger)
		logger.Error().Caller().AnErr("service.UserInfo", err).Send()
		// 从数据库查询
		result := svc.db.Model(model.User{}).Joins("left join user_ex on user.id = user_ex.user_id").
			Select("user.username, user.email, user.password, user.avatar, user.gender, user.root,user.introduce,user.created_at,"+
				"user_ex.works_nums,user_ex.work_day, user_ex.words,user_ex.fans_nums").
			First(userCache, "user.id=?", id)
		err = CheckError(result, Select_OP)
		if err != nil {
			return
		}
		// 存入缓存
		userCache.HMSetAll()
	}
	// 邮件是否展示
	currentUserId, exists := svc.ctx.Get("me.id")
	if !exists || currentUserId.(int64) != id {
		var userSt model.UserSt
		err = svc.db.Select("show_email").First(&userSt, "user_id=?", id).Error
		if err != nil {
			return
		}
		if !userSt.ShowEmail {
			userCache.Email = ""
		}
	}
	resp = UserResponse{
		ID:        userCache.ID,
		Username:  userCache.Username,
		Email:     userCache.Email,
		Root:      userCache.Root,
		Avatar:    userCache.Avatar,
		Gender:    userCache.Gender,
		Introduce: userCache.Introduce,
		WorksNums: userCache.WorksNums,
		WorkDay:   userCache.WorkDay,
		Words:     userCache.Words,
		FansNums:  userCache.FansNums,
		CreatedAt: userCache.CreatedAt,
	}
	return
}

// CurrentUserSetting: 当前用户偏好设置
func (svc Service) CurrentUserSetting() (st UserSettingResponse, err error) {
	defer errwrap.Wrap(&err, "service.CurrentUserSetting")

	// 获取用户偏好
	userId := svc.ctx.GetInt64("me.id")
	err = cache.New(st.TableName(), userId).Cache(svc.ctx.Request.Context(), &st, func() error {
		result := svc.db.First(&st, "user_id=?", userId)
		return CheckError(result, Select_OP)
	})
	return
}

// UpdateUserInfo: 修改用户信息
func (svc Service) UpdateUserInfo(req UserUpdateRequest) (err error) {
	defer errwrap.Wrap(&err, "service.UpdateUserInfo")

	err = svc.db.Transaction(func(tx *gorm.DB) error {
		id := svc.ctx.GetInt64("me.id")
		// 构造数据
		update := make(map[string]interface{})
		if req.Username != "" {
			update["username"] = req.Username
		}
		if req.Gender != nil {
			update["gender"] = *req.Gender
		}
		if req.Avatar != "" {
			update["avatar"] = req.Avatar
		}
		if req.Introduce != "" {
			update["introduce"] = req.Introduce
		}
		// 修改数据库信息
		result := tx.Model(model.User{ID: id}).Updates(update)
		err := CheckError(result, Update_OP)
		if err != nil {
			return err
		}
		// 修改redis信息
		userCache := &cache.UserCache{ID: id}
		return userCache.HMSetField(update)
	})
	return
}

// UpdateUserEmail: 修改用户邮箱
func (svc Service) UpdateUserEmail(req UserUpdateEmailRequest) (err error) {
	defer errwrap.Wrap(&err, "service.UpdateUserInfo")
	// 构造数据
	err = svc.db.Transaction(func(tx *gorm.DB) error {
		id := svc.ctx.GetInt64("me.id")
		var user model.User
		err := tx.Select("password").First(&user).Error
		if err != nil {
			return err
		}
		// 校验密码
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return errcode.ErrorUserPassword
		}
		if err != nil {
			return err
		}
		// 修改数据库信息
		result := tx.Model(model.User{ID: id}).Update("email", req.Email)
		err = CheckError(result, Update_OP)
		if err != nil {
			return err
		}
		// 修改redis信息
		userCache := &cache.UserCache{ID: id}
		return userCache.HMSetField(map[string]interface{}{"email": req.Email})
	})
	return
}

// UpdateUserPassword: 修改用户密码
func (svc Service) UpdateUserPassword(req UserUpdatePassRequest) (err error) {
	defer errwrap.Wrap(&err, "service.UpdateUserInfo")
	// 构造数据
	err = svc.db.Transaction(func(tx *gorm.DB) error {
		id := svc.ctx.GetInt64("me.id")
		var user model.User
		err := tx.Select("password").First(&user).Error
		if err != nil {
			return err
		}
		// 校验密码
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return errcode.ErrorUserPassword
		}
		if err != nil {
			return err
		}
		// 修改数据库信息
		pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			return err
		}
		result := tx.Model(model.User{ID: id}).Update("password", string(pass))
		err = CheckError(result, Update_OP)
		if err != nil {
			return err
		}
		// 修改redis信息
		userCache := &cache.UserCache{ID: id}
		return userCache.HMSetField(map[string]interface{}{"password": string(pass)})
	})
	return
}

// UpdateUserSetting: 修改当前用户偏好设置
func (svc Service) UpdateUserSetting(req UserSettingRequest) (err error) {
	defer errwrap.Wrap(&err, "service.UpdateUserSetting")

	err = svc.db.Transaction(func(tx *gorm.DB) error {
		userId := svc.ctx.GetInt64("me.id")
		result := tx.Model(model.UserSt{}).Updates(map[string]interface{}{
			"show_email":         req.ShowEmail,
			"disable_search":     req.DisableSearch,
			"show_adult":         req.ShowAdult,
			"hidden_grade":       req.HiddenGrade,
			"hidden_tag":         req.HiddenTag,
			"subscription_email": req.SubscriptionEmail,
			"topic_email":        req.TopicEmail,
			"comment_email":      req.CommentEmail,
			"system_email":       req.SystemEmail,
		})
		err = CheckError(result, Update_OP)
		if err != nil {
			return err
		}
		return cache.New(model.UserSt{}.TableName(), userId).HMSet(req)
	})
	return
}
