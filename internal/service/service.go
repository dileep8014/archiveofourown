package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/pkg/runner"
)

type Service struct {
	tx *runner.Runner
}

func NewService(c *gin.Context) *Service {
	return &Service{tx: runner.NewRunner(c)}
}

func (svc *Service) Finish() {
	svc.tx.Close()
}
