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
	ErrorCollegeWorks = NewError(20040001, "查询书单内作品列表失败")
)
