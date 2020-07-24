package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"github.com/shyptr/sqlex"
	"net/http"
)

func Tx() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.Value("logger").(zerolog.Logger)
		tx, err := global.DBEngine.Begin()
		if err != nil {
			logger.Error().AnErr("middleware.tx.begin", err).Send()
			res := app.NewResponse(c)
			res.ToErrorResponse(errcode.ServerError)
			c.Abort()
			return
		}
		c.Set("tx", sqlex.BaseRunner(tx))
		defer func() {
			if c.Writer.Status() == http.StatusOK {
				err := tx.Commit()
				if err != nil {
					logger.Error().AnErr("middleware.tx.commit", err).Send()
					tx.Rollback()
				}
				return
			}
			tx.Rollback()
		}()
		c.Next()
	}
}
