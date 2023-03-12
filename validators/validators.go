package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Trans ut.Translator
)

func ValidatorInit(locale string) error {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//国际化语言
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT)
		var found bool
		Trans, found = uni.GetTranslator(locale)
		if !found {
			fmt.Println("未找到对应语种")
		}
		//注册翻译
		switch locale {
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(validate, Trans)
		case "en":
			_ = en_translations.RegisterDefaultTranslations(validate, Trans)
		default:
			_ = zh_translations.RegisterDefaultTranslations(validate, Trans)
		}
		return nil
	} else {
		return errors.New("类型断言失败")
	}
}
func Translate(err error) gin.H {
	var result = gin.H{}
	validationErrors := err.(validator.ValidationErrors)
	for _, v := range validationErrors {
		result[v.Field()] = v.Translate(Trans)
	}
	return result
}
