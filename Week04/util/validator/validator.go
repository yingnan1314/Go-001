package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

func Validate(req interface{}) error {
	validate := validator.New()
	err := validate.Struct(req)

	cn := zh.New()
	uni := ut.New(cn, cn)
	translator, found := uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("not found")
	}
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			return err
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				return errors.New(err.Translate(translator))
			}
		}

	}
	return nil
}

