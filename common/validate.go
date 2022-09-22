/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-05-16 14:38:15
 * @LastEditors: Wynters
 * @LastEditTime: 2021-05-16 15:56:17
 */
package common

import (
	"errors"

	"github.com/go-playground/form"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate(s interface{}) error {
	zh := zh.New()
	trans, _ := ut.New(zh, zh).GetTranslator("zh")
	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(s)
	errStr := ""

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errStr = e.Translate(trans)
			return errors.New(errStr)
		}
	}
	return nil
}

func FormDecode(v interface{}, postValues map[string][]string) error {
	var decoder *form.Decoder
	decoder = form.NewDecoder()
	err := decoder.Decode(&v, postValues)
	if err != nil {
		return err
	}
	return nil
}
