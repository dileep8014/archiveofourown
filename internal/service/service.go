package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/gormplugins"
)

type Service struct {
	ctx *gin.Context
	db  *gorm.DB
}

func NewService(c *gin.Context) *Service {
	db := gormplugins.SetSpanToGorm(c.Request.Context(), global.Engine)
	name, exists := c.Get("me.name")
	if exists {
		db.Set("me.name", name)
	}
	return &Service{ctx: c, db: db}
}
