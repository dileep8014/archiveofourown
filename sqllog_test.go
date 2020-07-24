package main

import (
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/pkg/logger"
	"github.com/shyptr/archiveofourown/pkg/runner"
	"github.com/shyptr/archiveofourown/pkg/setting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArticle_Get(t *testing.T) {
	st, err := setting.NewSetting()
	assert.NoError(t, err)
	global.SetupSetting()
	logger.SetupLogger()
	err = global.SetupDBEngine(st)
	assert.NoError(t, err)
	article := model.Article{ID: 1}
	tx, _ := global.DBEngine.Begin()
	article, err = article.Get(&runner.Runner{Tx: tx, Logger: logger.Get()})
	if err != nil {
		tx.Rollback()
	}else {
		tx.Commit()
	}
	assert.NoError(t, err)
}
