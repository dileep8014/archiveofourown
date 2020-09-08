package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shyptr/archiveofourown/global"
	"strconv"
	"strings"
)

type ValidError struct {
	Key string
	Msg string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Msg
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func ShouldParamConvertInt(c *gin.Context, key string) (int, error) {
	v := c.Param(key)
	if v == "" {
		return 0, fmt.Errorf("uri 参数 %s 为空", key)
	}
	return strconv.Atoi(v)
}

func BindAndValid(c *gin.Context, v interface{}, bind func(obj interface{}) error) (bool, ValidErrors) {
	var errs ValidErrors
	err := bind(v)
	if err != nil {
		trans := global.Translations()
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return true, nil
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key: key,
				Msg: value,
			})
		}
		return false, errs
	}
	return true, nil
}
