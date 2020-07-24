package errcode

var (
	ErrorCreateCategoryFail = NewError(20010001, "创建分类失败")
	ErrorUpdateCategoryFail = NewError(20010002, "更新分类失败")
	ErrorListCategoryFail   = NewError(20010003, "查询所有分类失败")
	ErrorDeleteCategoryFail = NewError(20010004, "删除分类失败")
	ErrorGetUserFail        = NewError(20010005, "查找用户失败")
)
