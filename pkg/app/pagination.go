package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/convert"
)

func GetPage(c *gin.Context) (size, offset int) {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}
	size = convert.StrTo(c.Query("page_size")).MustInt()
	if size <= 0 {
		size = 1
	}
	if size > global.AppSetting.MaxPageSize {
		size = global.AppSetting.MaxPageSize
	}

	offset = (page - 1) * size
	return
}
