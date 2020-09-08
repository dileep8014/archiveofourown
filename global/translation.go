package global

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

var zh_trans ut.Translator

func InitValidate() {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh2.RegisterDefaultTranslations(v, trans)
	}
	zh_trans = trans
}

func Translations() ut.Translator {
	return zh_trans
}
