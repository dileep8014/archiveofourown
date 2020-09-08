package errcode

// calendar error
var (
	ErrorListCalendar = NewError(20010001, "查询日历列表失败")
)

// category error
var (
	ErrorCreateCategoryFail = NewError(20020001, "创建分类失败")
	ErrorUpdateCategoryFail = NewError(20020002, "更新分类失败")
	ErrorListCategoryFail   = NewError(20020003, "查询所有分类失败")
	ErrorDeleteCategoryFail = NewError(20020004, "删除分类失败")
)

// chapter error
var (
	ErrorGetChapter        = NewError(20030001, "查询章节内容失败")
	ErrorGetHistoryChapter = NewError(20030002, "查询章节历史版本失败")
	ErrorNewChapter        = NewError(20030003, "新建章节失败")
	ErrorSaveChapter       = NewError(20030004, "保存章节失败")
	ErrorPublishChapter    = NewError(20030005, "章节发布失败")
	ErrorUpdateChapter     = NewError(20030006, "修改分卷信息失败")
	ErrorDeleteChapter     = NewError(20030007, "删除章节失败")
	ErrorLockChapter       = NewError(20030008, "锁住章节失败")
	ErrorUnLockChapter     = NewError(20030009, "解锁章节失败")
)

// college error
var (
	ErrorCollegeWorks   = NewError(20040001, "查询书单内作品列表失败")
	ErrorCollegeList    = NewError(20040002, "查询书单列表失败")
	ErrorCreateCollege  = NewError(20040003, "创建书单失败")
	ErrorCollegeAddWork = NewError(20040004, "书单添加作品失败")
	ErrorCollegeUpdate  = NewError(20040005, "书单更新信息失败")
	ErrorCollegeDelete  = NewError(20040006, "书单删除失败")
)

// user error
var (
	ErrorEmailExist        = NewError(20050001, "邮箱已被注册")
	ErrorRegisterUser      = NewError(20050002, "注册失败")
	ErrorIdentifyUser      = NewError(20050003, "注册验证失败")
	ErrorCreateUser        = NewError(20050004, "创建用户账户失败")
	ErrorUsernameExist     = NewError(20050005, "用户名已被使用")
	ErrorUserNotExist      = NewError(20050006, "用户不存在")
	ErrorUserPassword      = NewError(20050007, "密码错误")
	ErrorUserLogin         = NewError(20050008, "登录失败")
	ErrorUserInfo          = NewError(20050009, "获取用户信息失败")
	ErrorUserSetting       = NewError(20050010, "获取用户偏好设置失败")
	ErrorUserUpdate        = NewError(20050011, "修改用户信息失败")
	ErrorUserUpdateSetting = NewError(20050012, "修改用户偏好设置失败")
)

var ErrorFileUpload = NewError(20060001, "文件上传失败")
