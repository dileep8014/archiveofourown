package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/tracer"
	"gorm.io/gorm"
)

type Service struct {
	ctx *gin.Context
	db  *gorm.DB
}

func NewService(c *gin.Context) *Service {
	return &Service{ctx: c, db: tracer.SetSpanToGorm(c, global.Engine)}
}
